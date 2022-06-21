package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
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

func main() {
	// Logger section
	log.SetFormatter(&log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime: "@timestamp",
			log.FieldKeyMsg:  "message",
		},
	})
	log.SetLevel(log.TraceLevel)

	psql := db.Connect(os.Getenv("DATABASE_URI"))
	siteRepo := site.NewRepository(psql)

	server := gin.New()
	server.GET("/", api.RootHandler())

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
