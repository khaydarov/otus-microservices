package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"users/api"
	"users/internal/db"
	"users/pkg/session"
	"users/pkg/user"
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

	// Persistence
	psql := db.Connect(os.Getenv("DATABASE_URI"))
	userRepo := user.NewRepository(psql)
	sessionRepo := session.NewRepository(psql)

	server := gin.New()
	server.GET("/", api.RootHandler())
	server.POST("/signup", api.SignUpHandler(userRepo))
	server.POST("/login", api.LoginHandler(userRepo, sessionRepo))

	err := server.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("Server is not started: %s", err)
	}
}
