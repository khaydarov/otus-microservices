package broker

import (
	"github.com/segmentio/kafka-go"
	"os"
)

func NewKafkaReader(topic string) *kafka.Reader {
	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{os.Getenv("KAFKA_HOST")},
		Topic: topic,
		GroupID: "consumer-group-1",
	})

	return kafkaReader
}
