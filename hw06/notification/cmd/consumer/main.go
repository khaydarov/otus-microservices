package main

import (
	"context"
	"encoding/json"
	"github.com/joho/godotenv"
	"hw06/notification/broker"
	"hw06/notification/database"
	"hw06/notification/pkg/notification"
	"log"
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

	reader := broker.NewKafkaReader("notifications")

	ctx := context.Background()
	log.Println("Consuming...")
	for {
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))

		var newNotification notification.Notification
		err = json.Unmarshal(m.Value, &newNotification)
		if err != nil {
			continue
		}

		newNotification.ID = notification.NewID()
		repository := notification.NewPsqlRepository(connection)
		err = repository.Store(newNotification)
		if err != nil {
			log.Printf("error: %s\n", err)
		}
	}
}


