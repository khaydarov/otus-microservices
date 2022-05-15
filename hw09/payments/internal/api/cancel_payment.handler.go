package api

import (
	"github.com/gin-gonic/gin"
	"hw09/payments/internal/service"
	"net/http"
)

// CancelPaymentHandler handles request to cancel payment
func CancelPaymentHandler(service service.PaymentService) func (c *gin.Context) {
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

		err := service.DeletePayment(body.OrderID)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": err.Error(),
				"data": gin.H{},
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "",
				"data": gin.H{},
			})
		}
	}
}
