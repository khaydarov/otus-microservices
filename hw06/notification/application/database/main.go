package database

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
	"os"
)

var connection *pgx.Conn

func Init() {
	var err error
	connection, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URI"))
	if err != nil {
		log.Fatalf("DB connection error: %s", err)
	}
}

func GetConnection() *pgx.Conn {
	return connection
}
