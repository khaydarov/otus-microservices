package api

import (
	"billing/pkg/account"
	"billing/pkg/legder"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BalanceHandler(accountRepo account.Repository, ledger legder.Ledger) func(c *gin.Context) {
	return func(c *gin.Context) {
		userID, ok := c.Get("UserID")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "invalid token",
			})

			return
		}

		userAccount, err := accountRepo.GetByUserID(userID.(string))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "user account not found",
			})

			return
		}

		balance := ledger.GetAccountBalance(userAccount)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "",
			"data": gin.H{
				"balance": balance,
			},
		})
	}
}
