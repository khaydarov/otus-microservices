package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"hw06/order/controllers"
	"hw06/order/internal/order"
	"hw06/order/middlewares"
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
	connection, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("DB connection error: %s", err)
	}
	defer connection.Close(context.Background())

	r := gin.New()
	r.GET("/", func (c *gin.Context) {
		c.JSON(200, "Hello to order service!")
	})

	r.Use(middlewares.AuthMiddleware())
	r.POST("/orders", controllers.CreateOrder(order.NewPsqlOrderRepository(connection)))
	err = r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("Server is not started: %s", err)
	}
}