package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"hw09/orders/internal/api"
	"hw09/orders/internal/db"
	"hw09/orders/internal/order"
	"hw09/orders/internal/tracer"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	ctx := context.Background()

	// Logger section
	log.SetFormatter(&log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime:  "@timestamp",
			log.FieldKeyMsg:   "message",
		},
	})
	log.SetLevel(log.TraceLevel)

	// Tracer section
	provider, err := tracer.NewProvider(tracer.ProviderConfig{
		JaegerEndpoint: os.Getenv("TRACER_PROVIDER"),
		ServiceName: "order-service",
		Environment: "production",
	})

	if err != nil {
		log.Fatal("Error creating tracing provider")
	}
	defer provider.Close(ctx)

	psql := db.Connect(os.Getenv("DATABASE_URI"))
	orderRepository := order.NewRepository(psql)

	server := gin.New()
	server.Use(otelgin.Middleware("payment-service"))
	server.GET("/", api.RootHandler())
	server.POST("/", api.CreateOrderHandler(orderRepository))

	err = server.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("Server is not started: %s", err)
	}
}

