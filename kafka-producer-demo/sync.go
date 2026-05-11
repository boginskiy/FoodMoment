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
	deliveryChan := make(chan kafka.Event)
	topic := "my-topic"
	mess := "Hello, Kafka!"

	// 4. Синхронная отправка сообщения.
	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny, // Автоматический выбор партиции.
		},
		Value: []byte(mess),
	}, deliveryChan)

	if err != nil {
		fmt.Printf("produce failed: %s\n", err)
		return
	}

	// 5. Ожидание подтверждения
	event := <-deliveryChan
	msg := event.(*kafka.Message)

	if msg.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", msg.TopicPartition.Error)
	} else {
		fmt.Printf("%s - [%d] at offset %v\n",
			*msg.TopicPartition.Topic,
			msg.TopicPartition.Partition,
			msg.TopicPartition.Offset)
	}

	close(deliveryChan)
}

//
// [Running] go run "/home/ali/dev/FoodMoment/kafka-producer-demo/main.go"
// %3|1778365131.250|FAIL|my-producer#producer-1| [thrd:localhost:9092/bootstrap]: localhost:9092/bootstrap: Connect to ipv4#127.0.0.1:9092 failed: Connection refused (after 0ms in state CONNECT)
// %3|1778365132.250|FAIL|my-producer#producer-1| [thrd:localhost:9092/bootstrap]: localhost:9092/bootstrap: Connect to ipv4#127.0.0.1:9092 failed: Connection refused (after 0ms in state CONNECT, 1 identical error(s) suppressed)
// %3|1778365162.251|FAIL|my-producer#producer-1| [thrd:localhost:9092/bootstrap]: localhost:9092/bootstrap: Connect to ipv4#127.0.0.1:9092 failed: Connection refused (after 0ms in state CONNECT, 30 identical error(s) suppressed)
