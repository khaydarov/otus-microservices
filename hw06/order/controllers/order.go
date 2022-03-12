package controllers

import (
	"github.com/gin-gonic/gin"
	"hw06/order/internal"
	"hw06/order/internal/order"
	"net/http"
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
			c.JSON(401, "Not authorized")
			return
		}

		var data createOrderRequestData
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"error": err.Error(),
			})

			return
		}

		user := credentials.(internal.User)
		newOrder := order.NewOrder(user.ID, data.Title, data.Price)
		err := orderRepository.Store(newOrder)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusInternalServerError,
				"error": err.Error(),
			})

			return
		}

		c.JSON(200, newOrder)
	}
}
