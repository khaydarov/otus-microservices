package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"hw09/inventory/internal/api"
	"hw09/inventory/internal/db"
	"hw09/inventory/internal/goods"
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
	goodsRepository := goods.NewRepository(psql)

	server := gin.Default()
	server.POST("/reserveGoods", api.ReserveGoodsHandler(goodsRepository))
	server.POST("/cancelGoodsReservation", api.CancelGoodsReservationHandler(goodsRepository))

	err := server.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("Server is not started: %s", err)
	}
}