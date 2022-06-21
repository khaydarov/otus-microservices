package main

import (
	"billing/api"
	"billing/internal/db"
	"billing/pkg/account"
	"billing/pkg/legder"
	"billing/pkg/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Logger section
	log.SetFormatter(&log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime: "@timestamp",
			log.FieldKeyMsg:  "message",
		},
	})
	log.SetLevel(log.TraceLevel)

	psql := db.Connect(os.Getenv("DATABASE_URI"))
	accountRepo := account.NewRepository(psql)
	ledger := legder.NewLedger(psql)

	server := gin.New()
	server.GET("/", api.RootHandler())

	publicApi := server.Group("/billing").Use(middleware.Auth())
	{
		publicApi.POST("/deposit", api.DepositHandler(accountRepo, ledger))
		publicApi.POST("/withdraw", api.WithdrawHandler(accountRepo, ledger))
		publicApi.GET("/balance", api.BalanceHandler(accountRepo, ledger))
	}

	internalApi := server.Group("/internal/billing")
	{
		internalApi.POST("/createAccount", api.CreateAccountHandler(accountRepo))
		internalApi.POST("/makePayment", api.MakePaymentHandler(accountRepo, ledger))
	}

	err := server.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("Server is not started: %s", err)
	}
}
