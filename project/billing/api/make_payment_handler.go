package api

import (
	"billing/pkg/account"
	"billing/pkg/legder"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func MakePaymentHandler(accountRepo account.Repository, ledger legder.Ledger) func(c *gin.Context) {
	type Body struct {
		Amount int    `json:"amount"`
		UserID string `json:"userID"`
	}

	return func(c *gin.Context) {
		var body Body
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": fmt.Sprintf("bad request: %s", err.Error()),
			})

			return
		}

		userAccount, err := accountRepo.GetByUserID(body.UserID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "user account not found",
			})

			return
		}

		revenue, err := accountRepo.GetByID(account.WithValue(os.Getenv("REVENUE_ID")))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "internal billing error",
			})

			return
		}

		transaction := legder.NewTransaction("Payment for advert")
		transaction.AddEntry(userAccount, 1, body.Amount)
		transaction.AddEntry(revenue, 2, body.Amount)

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
