package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// 1. Конфигурация Producer
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "my-producer",
		"acks":              "all", // Подтверждение от всех реплик [citation:6]
	}

	// 2. Создание Producer
	producer, err := kafka.NewProducer(config)
	if err != nil {
		fmt.Printf("failed to create producer: %s\n", err)
		os.Exit(1)
	}

	defer producer.Close() // Вот где мы закрываем продюсера!

	// 3. Канал для синхронных отчетов.
}
