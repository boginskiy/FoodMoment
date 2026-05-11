package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	config := &kafka.ConfigMap{
		"bootstrap.servers":            "kafka1:9092,kafka2:9092,kafka3:9092", // Минимум 2-3 брокера для отказоустойчивости
		"client.id":                    "LoggerWriter",                        // Логическое имя продюсера, которое отправляется в каждом запросе к брокерам
		"acks":                         "1",                                   // Компромисс. Только лидер партиции подтверждает получение данных. Реплики могут не успеть получить данные
		"go.produce.channel.size":      1000,                                  // Оптимизация производительности
		"queue.buffering.max.messages": 1000,                                  // Оптимизация производительности
	}

	//
	producer, err := kafka.NewProducer(config)
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	// Обработка отчетов о доставке в отдельной горутине
	go func() {
		for event := range producer.Events() {
			msg, ok := event.(*kafka.Message)
			if !ok {
				continue
			}

			if msg.TopicPartition.Error != nil {
				fmt.Println("failed to deliver: %v\n", msg.TopicPartition.Error)
				continue
			}

			// Result
			fmt.Println(
				*msg.TopicPartition.Topic,
				msg.TopicPartition.Partition,
				msg.TopicPartition.Offset,
			)
		}
	}()

	//
	topic := "async-topic"
	messages := []string{"first", "second", "third", "fourth", "fifth"}

	// Асинхронная отправка — не блокирует выполнение
	for _, word := range messages {
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil) // nil = нет канала для отчета, только через Events()

		fmt.Printf("Queued: %s\n", word)
	}

	// Graceful shutdown
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	<-sigchan

	// Ждем завершения всех отправок
	producer.Flush(15 * 1000) // 15 секунд
}
