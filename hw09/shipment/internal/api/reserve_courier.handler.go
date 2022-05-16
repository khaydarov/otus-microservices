package api

import (
	"github.com/gin-gonic/gin"
	"hw09/shipment/internal/courier"
	"net/http"
)

// ReserveCourierHandler handles request to reserve courier
func ReserveCourierHandler(repository courier.Repository) func (c *gin.Context) {
	// Request body structure
	type Body struct {
		OrderID string `json:"order_id"`
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

		freeCourier, err := repository.GetFreeCourier()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": err.Error(),
				"data": gin.H{},
			})

			return
		}

		err = repository.Reserve(freeCourier, body.OrderID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"data": gin.H{
					"success": false,
					"message": "could not reserve courier",
					"data": gin.H{
						"order_id": body.OrderID,
					},
				},
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"success": false,
				"message": "",
				"data": gin.H{
					"order_id": body.OrderID,
				},
			},
		})
	}
}
