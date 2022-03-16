package main

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"hw06/billing/account"
	"log"
	"os"
)

var postgresConnection *pgx.Conn
var kafkaReader *kafka.Reader

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	initDb()
	defer postgresConnection.Close(context.Background())

	initKafka()
	defer kafkaReader.Close()
}

func initDb() {
	var err error
	postgresConnection, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("DB connection error: %s", err)
	}
}

func initKafka() {
	kafkaReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{os.Getenv("KAFKA_HOST")},
		Topic: os.Getenv("KAFKA_TOPIC"),
		GroupID: "consumer-group-1",
	})

	ctx := context.Background()
	log.Println("Consuming...")
	for {
		m, err := kafkaReader.ReadMessage(ctx)
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))

		var userCreated UserCreated
		err = json.Unmarshal(m.Value, &userCreated)
		if err != nil {
			log.Printf("Error: %s\n", err)
			continue
		}

		repository := account.NewRepository(postgresConnection)
		newAccount := account.NewAccount(userCreated.ID)
		err = repository.Store(newAccount)
		if err != nil {
			log.Printf("Error: %s\n", err)
		}
	}
}

type UserCreated struct {
	ID 			string `json:"ID"`
	FirstName 	string `json:"FirstName"`
	LastName	string `json:"LastName"`
}
