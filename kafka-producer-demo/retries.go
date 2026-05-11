package main

import (
	"context"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type RetryableProducer struct {
	producer    *kafka.Producer
	maxRetries  int
	backoffBase time.Duration
}

func NewRetryableProducer(cfg *kafka.ConfigMap, maxRetries int, backoff time.Duration) (*RetryableProducer, error) {
	p, err := kafka.NewProducer(cfg)
	if err != nil {
		return nil, err
	}
	return &RetryableProducer{
		producer:    p,
		maxRetries:  maxRetries,
		backoffBase: backoff,
	}, nil
}

func (rp *RetryableProducer) ProduceWithRetry(ctx context.Context, msg *kafka.Message) error {
	var lastErr error

	for attempt := 0; attempt <= rp.maxRetries; attempt++ {
		// Канал для отчета
		deliveryChan := make(chan kafka.Event)

		err := rp.producer.Produce(msg, deliveryChan)
		if err != nil {
			lastErr = err
			if attempt < rp.maxRetries {
				rp.backoff(attempt)
				continue
			}
			return fmt.Errorf("produce failed after %d attempts: %w", attempt, err)
		}

		// Ожидание результата
		select {
		case e := <-deliveryChan:
			kafkaMsg := e.(*kafka.Message)
			if kafkaMsg.TopicPartition.Error != nil {
				lastErr = kafkaMsg.TopicPartition.Error
				if attempt < rp.maxRetries {
					rp.backoff(attempt)
					continue
				}
				return fmt.Errorf("delivery failed after %d attempts: %w", attempt, lastErr)
			}
			return nil

		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return lastErr
}

func (rp *RetryableProducer) backoff(attempt int) {
	backoff := rp.backoffBase * time.Duration(1<<attempt) // Exponential: 100ms, 200ms, 400ms...
	if backoff > 10*time.Second {
		backoff = 10 * time.Second
	}
	time.Sleep(backoff)
}

func (rp *RetryableProducer) Close() {
	rp.producer.Close()
}

func main() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"acks":              "all",
	}

	rp, _ := NewRetryableProducer(config, 3, 100*time.Millisecond)
	defer rp.Close()

	go func() {
		for e := range rp.producer.Events() {
			if msg, ok := e.(*kafka.Message); ok && msg.TopicPartition.Error != nil {
				// Логируем, но retry уже обрабатывается в ProduceWithRetry
			}
		}
	}()

	topic := "retry-topic"
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic},
		Value:          []byte("retryable message"),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := rp.ProduceWithRetry(ctx, msg)
	if err != nil {
		fmt.Printf("Failed to send message: %v\n", err)
		// Здесь можно отправить в Dead Letter Queue [citation:1]
	} else {
		fmt.Println("Message sent successfully")
	}
}
