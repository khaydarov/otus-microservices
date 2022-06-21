package main

import (
	"adverts/api"
	"adverts/internal/db"
	"adverts/internal/middleware"
	"adverts/pkg/advert"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
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
	advertRepo := advert.NewRepository(psql)
	advertSelector := advert.NewAdvertSelector(psql)

	server := gin.New()
	server.GET("/", api.RootHandler())
	server.GET("/adverts/relevant", api.GetRelevantAdvertHandler(advertRepo, advertSelector))

	publicApi := server.Group("/").Use(middleware.Auth())
	{
		publicApi.POST("/adverts", api.PostAdvertHandler(advertRepo))
	}

	err := server.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("Server is not started: %s", err)
	}
}
