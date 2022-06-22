package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"users/api"
	"users/internal/db"
	"users/internal/middleware"
	"users/pkg/session"
	"users/pkg/user"
)

var psql *pgx.Conn

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func initDb() {
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
	initDb()

	prometheus := middleware.NewPrometheus("users", "web", "/metrics")

	userRepo := user.NewRepository(psql)
	sessionRepo := session.NewRepository(psql)

	server := gin.New()
	server.GET("/", api.RootHandler())
	server.GET("/health", health())

	server.Use(prometheus.HandleFunc())
	server.GET("/metrics", metrics())

	publicApi := server.Group("/")
	{
		publicApi.POST("/signup", api.SignUpHandler(userRepo))
		publicApi.POST("/login", api.LoginHandler(userRepo, sessionRepo))
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
