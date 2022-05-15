package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"hw09/payments/internal/api"
	"hw09/payments/internal/db"
	"hw09/payments/internal/service"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}

func main() {
	psql := db.Connect(os.Getenv("DATABASE_URI"))
	paymentsSvc := service.NewPaymentService(psql)

	server := gin.Default()
	server.POST("/makePayment", api.MakePaymentHandler(paymentsSvc))
	server.POST("/cancelPayment", api.CancelPaymentHandler(paymentsSvc))

	err := server.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("server start failed: %s", err)
	}
}