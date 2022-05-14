package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"hw09/shipment/internal/api"
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
	server := gin.Default()
	server.POST("/reserveCourier", api.ReserveCourierHandler())
	server.POST("/cancelCourierReservation", api.CancelCourierReservationHandler())

	err := server.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("Server is not started: %s", err)
	}
}