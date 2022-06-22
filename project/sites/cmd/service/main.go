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
	"sites/api"
	"sites/internal/db"
	"sites/internal/middleware"
	"sites/pkg/site"
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

	prometheus := middleware.NewPrometheus("sites", "web", "/metrics")

	siteRepo := site.NewRepository(psql)

	server := gin.New()
	server.GET("/", api.RootHandler())
	server.GET("/health", health())

	server.Use(prometheus.HandleFunc())
	server.GET(prometheus.MetricsPath, metrics())

	publicApi := server.Group("/").Use(middleware.Auth())
	{
		publicApi.GET("/sites", api.GetSitesHandler(siteRepo))
		publicApi.POST("/sites", api.PostSitesHandler(siteRepo))
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
