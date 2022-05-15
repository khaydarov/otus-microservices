package api

import (
	"github.com/gin-gonic/gin"
	"hw09/inventory/internal/goods"
	"net/http"
)

// CancelGoodsReservationHandler handles HTTP request to cancel goods
func CancelGoodsReservationHandler(repository goods.Repository) func (c *gin.Context) {
	type Body struct {
		OrderId string `json:"order_id"`
	}

	return func (c *gin.Context) {
		body := Body{}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
				"data": gin.H{},
			})

			return
		}

		err := repository.CancelReservation(body.OrderId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": err.Error(),
				"data": gin.H{},
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "",
			"data": gin.H{},
		})
	}
}
