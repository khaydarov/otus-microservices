package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"hw07/order/internal/order"
	"hw07/order/middlewares"
	"log"
	"net/http"
	"os"
)

type createOrderRequestData struct {
	Title string
	Price int
}

// CreateOrder creates new order
func CreateOrder(orderService *order.Service, kafkaWriter *kafka.Writer) func (c *gin.Context) {
	return func (c *gin.Context) {
		credentials, ok := c.Get("user")
		if !ok {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		idempotencyKey := c.GetHeader("IdempotencyKey")
		if idempotencyKey == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "idempotency key is empty",
			})

			return
		}

		var data createOrderRequestData
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		orderID, err := orderService.GetOrderIDByIdempotencyKey(idempotencyKey)
		if err == nil {
			c.JSON(http.StatusCreated, gin.H{
				"data": gin.H{
					"id": orderID,
				},
			})

			return
		}

		user := credentials.(middlewares.User)
		newOrder := order.NewOrder(user.ID, data.Title, data.Price)

		err = orderService.Store(idempotencyKey, newOrder)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		endpoint := fmt.Sprintf("%s/withdrawAccount",
			os.Getenv("BILLING_SERVICE"),
		)

		var jsonData = []byte(fmt.Sprintf(`{"amount": %d}`, newOrder.Price))

		request, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		request.Header.Set("Authorization", c.GetHeader("Authorization"))

		client := &http.Client{}
		response, _ := client.Do(request)
		defer response.Body.Close()

		kafkaWriter.Topic = "notifications"
		if response.StatusCode != http.StatusOK {
			c.JSON(response.StatusCode, gin.H{
				"error": "not enough money",
			})

			message := map[string]string{
				"text": "order was not created",
				"userId": user.ID,
			}

			v, _ := json.Marshal(message)
			kafkaWriter.WriteMessages(
				context.Background(),
				kafka.Message{
					Value: v,
				},
			)

			return
		}

		message := map[string]string{
			"text": "order was created",
			"userId": user.ID,
		}

		v, _ := json.Marshal(message)
		err = kafkaWriter.WriteMessages(
			context.Background(),
			kafka.Message{
				Value: v,
			},
		)
		log.Printf("kafka error: %s\n", err)

		// Send notification
		c.JSON(http.StatusCreated, gin.H{
			"data": gin.H{
				"id": newOrder.ID.Value,
			},
		})
	}
}