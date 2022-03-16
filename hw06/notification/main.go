package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"hw06/notification/database"
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
	database.Init()
	connection := database.GetConnection()
	defer connection.Close(context.Background())

	r := gin.New()
	r.GET("/", func (c *gin.Context) {
		c.JSON(200, "Hello to notification service!")
	})

	err := r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("Server is not started: %s", err)
	}
}