package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"hw09/shipment/internal/api"
	"hw09/shipment/internal/courier"
	"hw09/shipment/internal/db"
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
	psql := db.Connect(os.Getenv("DATABASE_URI"))

	courierRepository := courier.NewCourierRepository(psql)

	server.POST("/reserveCourier", api.ReserveCourierHandler(courierRepository))
	server.POST("/cancelCourierReservation", api.CancelCourierReservationHandler(courierRepository))

	err := server.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("Server is not started: %s", err)
	}
}