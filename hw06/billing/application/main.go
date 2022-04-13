package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"hw06/billing/account"
	"hw06/billing/middlewares"
	"log"
	"net/http"
	"os"
)

var (
	postgresConnection *pgx.Conn
)

func main() {
	initDb()
	defer postgresConnection.Close(context.Background())

	r := gin.Default()
	r.GET("/health", health())
	r.GET("/metrics", metrics())

	p := middlewares.NewPrometheus("billing", "http")
	r.Use(p.HandleFunc())
	r.GET("/", func (c *gin.Context) {
		c.JSON(200, "Hello to billing service!")
	})

	// Register modules
	r.Use(middlewares.AuthMiddleware())
	account.RegisterRoutes(r.Group(""), postgresConnection)

	err := r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("Server is not started: %s", err)
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func initDb() {
	var err error
	postgresConnection, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URI"))
	if err != nil {
		log.Fatalf("DB connection error: %s", err)
	}
}

func health() func (c *gin.Context) {
	return func (c *gin.Context) {
		err := postgresConnection.Ping(context.Background())
		if err == nil {
			c.JSON(http.StatusOK, "Healthy")

			return
		}

		c.JSON(http.StatusInternalServerError, "Unhealthy")
	}
}

func metrics() func (c *gin.Context) {
	h := promhttp.Handler()
	return func (c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
