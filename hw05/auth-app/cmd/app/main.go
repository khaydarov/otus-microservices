package main

import (
	"auth-app/cmd/app/server"
	"auth-app/cmd/app/server/http"
	"auth-app/repository"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connection, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("DB connection error: %s", err)
	}
	defer connection.Close(context.Background())

	httpServer := http.Server{
		SessionRepository: repository.NewSessionRepository(*connection),
	}

	httpServer.Run(server.Config{
		Port: os.Getenv("APP_PORT"),
	})
}