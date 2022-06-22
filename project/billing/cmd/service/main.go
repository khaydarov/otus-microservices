package main

import (
	"billing/api"
	"billing/internal/db"
	"billing/pkg/account"
	"billing/pkg/legder"
	"billing/pkg/middleware"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

var psql *pgx.Conn

func initDB() {
	psql = db.Connect(os.Getenv("DATABASE_URI"))
}

func main() {
	// Logger section
	log.SetFormatter(&log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime: "@timestamp",
			log.FieldKeyMsg:  "message",
		},
	})
	log.SetLevel(log.TraceLevel)
	initDB()

	prometheus := middleware.NewPrometheus("billing", "web", "/metrics")

	accountRepo := account.NewRepository(psql)
	ledger := legder.NewLedger(psql)

	server := gin.New()
	server.GET("/", api.RootHandler())
	server.GET("/health", health())

	server.Use(prometheus.HandleFunc())
	server.GET("/metrics", metrics())

	publicApi := server.Group("/billing").Use(middleware.Auth())
	{
		publicApi.POST("/deposit", api.DepositHandler(accountRepo, ledger))
		publicApi.POST("/withdraw", api.WithdrawHandler(accountRepo, ledger))
		publicApi.GET("/balance", api.BalanceHandler(accountRepo, ledger))
	}

	internalApi := server.Group("/internal/billing")
	{
		internalApi.POST("/createAccount", api.CreateAccountHandler(accountRepo))
		internalApi.POST("/makePayment", api.MakePaymentHandler(accountRepo, ledger))
	}

	err := server.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("Server is not started: %s", err)
	}
}

func health() func(c *gin.Context) {
	return func(c *gin.Context) {
		err := psql.Ping(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Unhealthy")

			return
		}

		c.JSON(http.StatusOK, "Healthy")
	}
}

func metrics() func(c *gin.Context) {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
