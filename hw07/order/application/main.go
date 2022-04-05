package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"hw07/order/controllers"
	"hw07/order/internal/order"
	"hw07/order/middlewares"
	"log"
	"net/http"
	"os"
)

var (
	db *pgx.Conn
	kafkaWriter *kafka.Writer
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	initDB()
	defer db.Close(context.Background())

	initKafkaWriter()
	defer kafkaWriter.Close()

	r := gin.New()
	r.GET("/", func (c *gin.Context) {
		c.JSON(200, "Hello to order service!")
	})
	r.GET("/health", health())

	r.Use(middlewares.AuthMiddleware())
	r.POST("/", controllers.CreateOrder(
		order.NewService(db),
		kafkaWriter,
	))

	err := r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("Server is not started: %s", err)
	}
}

func initDB() {
	var err error
	ctx := context.Background()
	db, err = pgx.Connect(ctx, os.Getenv("DATABASE_URI"))

	if err != nil {
		log.Fatalf("DB connection error: %s", err)
	}
}

func initKafkaWriter() {
	kafkaWriter = &kafka.Writer{
		Addr: kafka.TCP(os.Getenv("KAFKA_HOST")),
		ErrorLogger: kafka.LoggerFunc(func (message string, args ...interface{}) {
			log.Println(message, args)
		}),
	}
}

func health() func (c *gin.Context) {
	return func (c *gin.Context) {
		err := db.Ping(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Unhealthy")

			return
		}

		c.JSON(http.StatusOK, "Healthy")
	}
}