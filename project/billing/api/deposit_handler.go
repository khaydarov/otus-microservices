package api

import (
	"billing/pkg/account"
	"billing/pkg/legder"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func DepositHandler(accountRepo account.Repository, ledger legder.Ledger) func(c *gin.Context) {
	type Body struct {
		Amount int `json:"amount"`
	}

	return func(c *gin.Context) {
		userID, ok := c.Get("UserID")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "invalid token",
			})

			return
		}

		var body Body
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "bad request",
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

		cashbook, err := accountRepo.GetByID(account.WithValue(os.Getenv("CASHBOOK_ID")))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "internal billing error",
			})

			return
		}

		transaction := ledger.NewTransaction("deposit")
		transaction.AddEntry(cashbook, 1, body.Amount)
		transaction.AddEntry(userAccount, 2, body.Amount)

		err = ledger.Commit(transaction)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		id := transaction.GetID()
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "transaction committed",
			"data": gin.H{
				"transactionID": id.GetValue(),
			},
		})
	}
}
