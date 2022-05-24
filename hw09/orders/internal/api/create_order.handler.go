package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"hw09/orders/internal/order"
	"hw09/orders/internal/saga"
	"hw09/orders/internal/service/inventory"
	"hw09/orders/internal/service/payments"
	"hw09/orders/internal/service/shipment"
	"hw09/orders/internal/tracer"
	"log"
	"net/http"
)

// CreateOrderHandler handles request to create order
func CreateOrderHandler(repository order.Repository) func (c *gin.Context) {
	type Good struct {
		ID 		int `json:"id"`
		Price 	int `json:"price"`
	}

	type Body struct {
		Goods []Good
	}

	return func (c *gin.Context) {
		ctx, span := tracer.NewSpan(context.Background(), "POST /order")
		defer span.End()

		body := Body{}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
				"data": gin.H{},
			})

			return
		}

		var goodIds []int
		amount := 0
		for _, good := range body.Goods {
			amount += good.Price
			goodIds = append(goodIds, good.ID)
		}

		span.AddEvent("create order")
		o := order.CreateOrder()
		err := repository.Store(ctx, o)

		if err != nil {
			span.AddEvent("order creation failed")
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
				"data": gin.H{},
			})

			return
		}

		span.AddEvent("order created")
		log.Println("order created!")

		s := saga.Saga{}
		s.SetName("order creation")
		s.AddStep(saga.Step{
			Name: "make payment",
			Func: func() error {
				ctx, span := tracer.NewSpan(ctx, "POST /makePayment")
				defer span.End()

				span.AddEvent("payments: start payment")
				err := payments.MakePayment(ctx, o.ID.GetValue(), amount)

				if err != nil {
					return err
				}

				span.AddEvent("payments: end payment")
				return nil
			},
			Compensation: func() error {
				_, span := tracer.NewSpan(ctx, "POST /makePayment")
				defer span.End()

				span.AddEvent("payments: cancel payment")

				err := payments.CancelPayment(o.ID.GetValue())
				if err != nil {
					return err
				}

				span.AddEvent("payments: canceled")

				return nil
			},
		})

		s.AddStep(saga.Step{
			Name: "reserve goods",
			Func: func() error {
				log.Println("inventory: start goods reservation")
				_, err := inventory.ReserveGoods(o.ID.GetValue(), goodIds)

				if err != nil {
					return err
				}

				log.Println("inventory: end goods reservation.")
				return nil
			},
			Compensation: func() error {
				log.Println("inventory: cancel goods reservation")

				err := inventory.CancelGoodsReservation(o.ID.GetValue())
				if err != nil {
					return err
				}

				return nil
			},
		})

		s.AddStep(saga.Step{
			Name: "reserve courier",
			Func: func() error {
				log.Println("shipment: start courier reservation")
				err := shipment.ReserveCourier(o.ID.GetValue())

				if err != nil {
					return err
				}

				log.Println("shipment: end courier reservation.")
				return nil
			},
			Compensation: func() error {
				log.Println("shipment: cancel courier reservation")

				err := shipment.CancelCourierReservation(o.ID.GetValue())
				if err != nil {
					return err
				}

				return nil
			},
		})

		coordinator := saga.NewCoordinator(s)
		err = coordinator.Commit()

		if err != nil {
			log.Println("order cancelled")

			repository.Delete(o)

			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
				"data": gin.H{},
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "",
			"data": gin.H{
				"order_id": o.ID.GetValue(),
			},
		})
	}
}