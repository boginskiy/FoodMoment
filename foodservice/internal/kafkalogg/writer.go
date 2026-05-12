package kafkalogg

import (
	"context"
	"errors"
	"time"

	"github.com/boginskiy/FoodMoment/foodservice/internal/logg"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaWrite struct {
	producer *kafka.Producer
	Logger   logg.Logger
}

func NewKafkaWrite(ctx context.Context, cfg Config, log logg.Logger) (*KafkaWrite, error) {
	// Config
	config := &kafka.ConfigMap{
		"bootstrap.servers":            cfg.Servers,
		"client.id":                    cfg.ClientID,
		"acks":                         cfg.Acks,
		"go.produce.channel.size":      cfg.ChannelSize,
		"queue.buffering.max.messages": cfg.MaxMessages,
	}
	// Producer
	p, err := kafka.NewProducer(config)
	if err != nil {
		return nil, err
	}

	tmp := &KafkaWrite{producer: p, Logger: log}

	go tmp.procResponse(ctx) // Обработка отчетов о доставке
	go tmp.doFlush(ctx)      // Отправка всех сообщений из буфера в сеть
	return tmp, nil
}

func (k *KafkaWrite) Close() error {
	if k.producer != nil {
		k.producer.Close()
		return nil
	}
	return errors.New("producer is not initialized")
}

func (k *KafkaWrite) Send(topic string, msg []byte) {
	k.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msg,
	}, nil)
}

func (k *KafkaWrite) doFlush(ctx context.Context) {
	<-ctx.Done()
	k.producer.Flush(workingTimeAfterDone * 1000)
}

func (k *KafkaWrite) procResponse(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			k.workingDoneWithTime(workingTimeAfterDone)
			return

		case event, ok := <-k.producer.Events():
			if !ok {
				return
			}
			msg, ok := event.(*kafka.Message)
			if !ok {
				continue
			}
			if msg.TopicPartition.Error != nil {
				k.Logger.Error("failed to check response",
					"error", msg.TopicPartition.Error)
				continue
			}
			// Result here, if need
			// *msg.TopicPartition.Topic,
			//  msg.TopicPartition.Partition,
			// 	msg.TopicPartition.Offset
		}
	}
}

func (k *KafkaWrite) workingDoneWithTime(sec int) {
	// Завершаем работу. Проверка оставшихся событий в заданном лимите времени.
	for {
		select {
		case event, ok := <-k.producer.Events():
			if !ok {
				return
			}
			msg, ok := event.(*kafka.Message)
			if !ok {
				continue
			}

			if msg.TopicPartition.Error != nil {
				k.Logger.Error("failed to check response",
					"error", msg.TopicPartition.Error)
				continue
			}

		case <-time.After(time.Duration(sec) * time.Second):
			return
		}
	}
}
