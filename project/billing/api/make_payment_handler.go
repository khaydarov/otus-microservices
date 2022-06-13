package api

import (
	"billing/pkg/account"
	"billing/pkg/transaction"
	"github.com/gin-gonic/gin"
)

func MakePaymentHandler(accountRepo account.Repository, transactionRepo transaction.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		transaction := transaction.NewTransaction("Оплата показов")
		transaction.AddEntry(account.Account{}, 1, 100)
		transaction.AddEntry(account.Account{}, 2, 100)

		transactionRepo.Commit(transaction)
	}
}
