package logg

import (
	"context"
	"log/slog"
)

const CapQueue = 100

type KafkaSend struct {
	buffCh   chan slog.Record
	Logger   Logger
	Preparer Preparer
	Writer   KafkaWriter
}

func NewKafkaSend(ctx context.Context, cfg KafkaConfig, log Logger) (*KafkaSend, error) {
	// Writer

	// TODO Создать Writer и далее делать sender, так скажем причесывать...

	tmp := &KafkaSend{
		buffCh:   make(chan slog.Record, CapQueue),
		Logger:   log,
		Preparer: &Prepar{},
	}

	go tmp.sendbroker(ctx)
	return tmp, nil
}

func (l *KafkaSend) Send(ctx context.Context, r slog.Record) {
	select {
	case <-ctx.Done():
		// Проверка наличия данных в канале.
		if len(l.buffCh) > 0 {
			// TODO...
			// Доработать тут . Сервер падает если
			// Если данные есть отправляем их в брокер.
			// Закрыть канал
		}

		close(l.buffCh)
		return

	case l.buffCh <- r:

	default:
		l.Logger.Info("buffer overflow when sending data to Kafka",
			"struct", "KafkaSend",
		)
	}
}

func (l *KafkaSend) sendbroker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			if len(l.buffCh) > 0 {
				// TODO...
			}
			return

		case v, ok := <-l.buffCh:
			if !ok {
				return
			}

			byteJSON, err := l.Preparer.SerializeToJSON(v)

			if err != nil {
				l.Logger.Error("failed to serialize json",
					"struct", "KafkaSend",
					"error", err)
				continue
			}

			// TODO... Готов byteJSON!
			writer.Send(ctx)
		}
	}
}

func (l *KafkaSend) Close() error {
	return nil
}
