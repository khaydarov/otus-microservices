package api

import (
	"billing/pkg/account"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAccountHandler(accountRepo account.Repository) func(c *gin.Context) {
	type Body struct {
		UserID string `json:"user_id"`
	}

	return func(c *gin.Context) {
		var body Body
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
				"data":    gin.H{},
			})

			return
		}

		newAccount := account.NewAccount(body.UserID)
		err := accountRepo.Store(newAccount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err,
				"data":    gin.H{},
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "",
			"data": gin.H{
				"account_id": newAccount.ID.GetValue(),
			},
		})
	}
}
