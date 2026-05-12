package kafkalogg

import (
	"context"
	"log/slog"

	"github.com/boginskiy/FoodMoment/foodservice/internal/logg"
)

const CapQueue = 100

type KafkaSend struct {
	buffCh   chan slog.Record
	Logger   logg.Logger
	Preparer Preparer
	Writer   Writer
	Cfg      Config
}

func NewKafkaSend(ctx context.Context, cfg Config, log logg.Logger) (*KafkaSend, error) {
	// Writer
	writer, err := NewKafkaWrite(ctx, cfg, log)
	if err != nil {
		return nil, err
	}

	tmp := &KafkaSend{
		buffCh:   make(chan slog.Record, CapQueue),
		Logger:   log,
		Writer:   writer,
		Preparer: &Prepar{},
		Cfg:      cfg,
	}

	go tmp.readBuffChan(ctx)
	return tmp, nil
}

func (l *KafkaSend) Send(ctx context.Context, r slog.Record) {
	defer func() {
		// Отлавливаем panic.
		if r := recover(); r != nil {
			l.Logger.Info("panic in a sending mess to buffer",
				"struct", "KafkaSend",
			)
		}
	}()

	select {
	case <-ctx.Done():
		return
	case l.buffCh <- r:
	default:
		l.Logger.Info("buffer overflow when sending data to Kafka",
			"struct", "KafkaSend",
		)
	}
}

func (l *KafkaSend) sendToBroker(record slog.Record) {
	byteJSON, err := l.Preparer.SerializeToJSON(record)

	if err != nil {
		l.Logger.Error("failed to serialize json",
			"struct", "KafkaSend",
			"error", err)
		return
	}
	// Send to Broker
	l.Writer.Send(l.Cfg.Topic, byteJSON)
}

func (l *KafkaSend) readBuffChan(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			if len(l.buffCh) > 0 {
				// Отправляем остатки сообщений из буфера
				for record := range l.buffCh {
					l.sendToBroker(record)
				}
			}
			return

		case record, ok := <-l.buffCh:
			if !ok {
				return
			}
			l.sendToBroker(record)
		}
	}
}

func (l *KafkaSend) Close() error {
	// Закрываем канал именно тут.
	defer close(l.buffCh)

	if err := l.Logger.Close(); err != nil {
		return err
	}

	if err := l.Writer.Close(); err != nil {
		return err
	}
	return nil
}
