package api

import (
	"github.com/gin-gonic/gin"
	"hw09/payments/internal/service"
	"net/http"
)

const limit = 1000

// MakePaymentHandler handles request to make payment
func MakePaymentHandler(service service.PaymentService) func (c *gin.Context) {
	// Request body structure
	type Body struct {
		OrderID string 	`json:"order_id"`
		Amount 	int 	`json:"amount"`
	}

	return func (c *gin.Context) {
		body := Body{}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success" : false,
				"message": err.Error(),
				"data": gin.H{},
			})

			return
		}

		if body.Amount > limit {
			c.JSON(http.StatusBadRequest, gin.H{
				"success" : false,
				"message": "not enough money",
				"data": gin.H{},
			})
		} else {
			err := service.StorePayment(body.OrderID, body.Amount)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success" : false,
					"message": err.Error(),
					"data": gin.H{},
				})

				return
			}

			c.JSON(http.StatusOK, gin.H{
				"success" : true,
				"message": "",
				"data": gin.H{},
			})
		}
	}
}