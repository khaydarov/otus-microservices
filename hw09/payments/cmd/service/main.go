package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"hw09/payments/internal/api"
	"hw09/payments/internal/db"
	"hw09/payments/internal/service"
	"hw09/payments/internal/tracer"
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
	ctx := context.Background()

	// Tracer section
	provider, err := tracer.NewProvider(tracer.ProviderConfig{
		JaegerEndpoint: os.Getenv("TRACER_PROVIDER"),
		ServiceName: "payment-service",
		Environment: "production",
	})

	if err != nil {
		log.Fatal("Error creating tracing provider")
	}
	defer provider.Close(ctx)

	psql := db.Connect(os.Getenv("DATABASE_URI"))
	paymentsSvc := service.NewPaymentService(psql)

	server := gin.Default()
	server.Use(otelgin.Middleware("payment-service"))
	server.POST("/makePayment", api.MakePaymentHandler(paymentsSvc))
	server.POST("/cancelPayment", api.CancelPaymentHandler(paymentsSvc))

	err = server.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("server start failed: %s", err)
	}
}