package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"hw09/orders/internal/api"
	"hw09/orders/internal/db"
	"hw09/orders/internal/order"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	psql := db.Connect(os.Getenv("DATABASE_URI"))
	orderRepository := order.NewRepository(psql)

	server := gin.Default()
	server.GET("/", api.RootHandler())
	server.POST("/", api.CreateOrderHandler(orderRepository))

	err := server.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("Server is not started: %s", err)
	}
}

