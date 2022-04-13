package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"hw06/notification/database"
	"hw06/notification/middlewares"
	"hw06/notification/pkg/notification"
	"log"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	database.Init()
	connection := database.GetConnection()
	defer connection.Close(context.Background())

	r := gin.New()
	r.Use(middlewares.AuthMiddleware())
	r.GET("/health", health())
	r.GET("/metrics", metrics())
	p := middlewares.NewPrometheus("notification", "http")
	r.Use(p.HandleFunc())
	r.GET("/", func (c *gin.Context) {
		credentials, ok := c.Get("user")
		if !ok {
			c.JSON(http.StatusUnauthorized, "")
			return
		}
		user := credentials.(middlewares.User)
		repository := notification.NewPsqlRepository(connection)
		notifications := repository.GetByUserID(user.ID)

		var response []map[string]string
		for _, n := range notifications {
			response = append(response, map[string]string{
				"id":   n.ID.Value,
				"text": n.Text,
			})
		}

		c.JSON(200, response)
	})

	err := r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("Server is not started: %s", err)
	}
}

func health() func (c *gin.Context) {
	return func (c *gin.Context) {
		connection := database.GetConnection()
		err := connection.Ping(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Unhealthy")

			return
		}

		c.JSON(http.StatusOK, "Healthy")
	}
}

func metrics() func (c *gin.Context) {
	h := promhttp.Handler()
	return func (c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}