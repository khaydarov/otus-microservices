package db

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
)

// Connect connects to postgreSQL and returns connection
func Connect(dsn string) *pgx.Conn {
	var err error
	ctx := context.Background()
	connection, err := pgx.Connect(ctx, dsn)

	if err != nil {
		log.Fatalf("DB connection error: %s", err)
	}

	return connection
}