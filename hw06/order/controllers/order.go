package controllers

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"hw06/order/internal/order"
	"hw06/order/middlewares"
	"net/http"
	"os"
)

type createOrderRequestData struct {
	Title string
	Price int
}

// CreateOrder creates new order
func CreateOrder(orderRepository order.Repository) func (c *gin.Context) {
	return func (c *gin.Context) {
		credentials, ok := c.Get("user")
		if !ok {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		var data createOrderRequestData
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		user := credentials.(middlewares.User)
		newOrder := order.NewOrder(user.ID, data.Title, data.Price)
		err := orderRepository.Store(newOrder)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		// TODO: refactor
		endpoint := fmt.Sprintf("%s/accounts/withdraw",
			os.Getenv("BILLING_SERVICE"),
		)

		var jsonData = []byte(fmt.Sprintf(`{"amount": %d}`, newOrder.Price))

		request, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		request.Header.Set("Authorization", c.GetHeader("Authorization"))

		client := &http.Client{}
		response, _ := client.Do(request)
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal error",
			})

			// Send notification
			return
		}

		// Send notification
		c.JSON(http.StatusCreated, gin.H{
			"data": gin.H{
				"id": newOrder.ID.Value,
			},
		})
	}
}