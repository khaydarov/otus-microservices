package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ReserveGoodsHandler handles HTTP request to reserve goods
func ReserveGoodsHandler() func (c *gin.Context) {
	type Body struct {
		OrderID string 	`json:"order_id"`
		GoodIds []int 	`json:"good_ids"`
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

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "",
			"data": gin.H{
				"order_id": body.OrderID,
				"good_ids": body.GoodIds,
			},
		})
	}
}
