package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"hw09/payments/internal/api"
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
	server := gin.Default()
	server.POST("/makePayment", api.MakePaymentHandler())
	server.POST("/cancelPayment", api.CancelPaymentHandler())

	err := server.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("server start failed: %s", err)
	}
}