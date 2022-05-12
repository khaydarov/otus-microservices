package api

import (
	"github.com/gin-gonic/gin"
	"hw09/orders/internal/saga"
	"log"
)

func CreateOrderHandler() func (c *gin.Context) {
	return func (c *gin.Context) {
		s := saga.Saga{}
		s.SetName("order creation")
		s.AddStep(saga.Step{
			Name: "withdraw money",
			Func: func() error {
				log.Println("withdrawing money")

				return nil
			},
			Compensation: func() error {
				log.Println("cancel withdrawal")

				return nil
			},
		})
		s.AddStep(saga.Step{
			Name: "reserve goods",
			Func: func() error {
				log.Println("reserving goods")

				return nil
			},
			Compensation: func() error {
				log.Println("cancel reservation")

				return nil
			},
		})

		coordinator := saga.NewCoordinator(s)
		err := coordinator.Commit()
		log.Println(err)
	}
}
