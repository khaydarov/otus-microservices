package api

import (
	"github.com/gin-gonic/gin"
	"hw09/shipment/internal/courier"
	"net/http"
)

// CancelCourierReservationHandler handles request to cancel courier reservation
func CancelCourierReservationHandler(repository courier.Repository) func (c *gin.Context) {
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

		err := repository.CancelReservation(body.OrderID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"data": gin.H{
					"success": false,
					"message": "could not cancel reservation",
					"data": gin.H{},
				},
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"success": true,
				"message": "",
				"data": gin.H{},
			},
		})
	}
}
