package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"hw06/user/infrastructure/metrics"
	"hw06/user/infrastructure/session"
	"hw06/user/infrastructure/user"
	"hw06/user/ui/web"
	"log"
	"net/http"
	"os"
)

var postgresConnection *pgx.Conn

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	initDb()
	defer postgresConnection.Close(context.Background())

	p := metrics.NewPrometheus("users", "http", "/metrics")
	r := gin.Default()
	r.GET("/health", health())
	r.Use(p.HandleFunc())
	r.GET(p.MetricsPath, prometheus())
	r.GET("/", func (c *gin.Context) {
		c.JSON(200, "Hello to user service!")
	})
	r.Use(gin.CustomRecovery(func (c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal error",
		})
	}))

	r.POST("/register", web.Register(
		user.NewPsqlRepository(postgresConnection)),
	)
	r.POST("/login", web.Login(
		user.NewPsqlRepository(postgresConnection),
		session.NewSessionRepository(postgresConnection),
	))
	r.POST("/auth", web.Auth(user.NewPsqlRepository(postgresConnection)))
	err := r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("Server is not started: %s", err)
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

func prometheus() func (c *gin.Context) {
	h := promhttp.Handler()

	return func (c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}