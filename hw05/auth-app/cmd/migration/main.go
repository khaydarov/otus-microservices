package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	m, err := migrate.New(
		"file://migrations",
		fmt.Sprintf("%s?sslmode=disable", os.Getenv("DATABASE_URL")))

	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}