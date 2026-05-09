package logg

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"
)

type ProduceLogg struct {
	buffCh   chan slog.Record
	capQueue int
	Logger   Logger
}

func NewProduceLogg(ctx context.Context, logg Logger, cap int) *ProduceLogg {
	tmp := &ProduceLogg{
		buffCh:   make(chan slog.Record, cap),
		capQueue: cap,
		Logger:   logg,
	}
	go tmp.sendbroker(ctx)
	return tmp
}

func (l *ProduceLogg) Send(ctx context.Context, r slog.Record) {
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
			"struct", "ProduceLogg",
		)
	}
}

func (l *ProduceLogg) Close() error {
	return nil
}

func (l *ProduceLogg) SerializeToJSON(r slog.Record) ([]byte, error) {
	viewLog := map[string]interface{}{
		"time":    r.Time.Format(time.RFC3339),
		"level":   r.Level.String(),
		"message": r.Message,
		"attrs":   make(map[string]interface{}),
	}

	attrs := make(map[string]interface{})
	r.Attrs(func(a slog.Attr) bool {
		attrs[a.Key] = a.Value.Any()
		return true
	})

	if len(attrs) > 0 {
		viewLog["attrs"] = attrs
	}

	return json.Marshal(viewLog)
}

func (l *ProduceLogg) sendbroker(ctx context.Context) {
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

			byteJSON, err := l.SerializeToJSON(v)

			if err != nil {
				l.Logger.Error("failed to serialize json",
					"struct", "ProduceLogg",
					"error", err)
				continue
			}

			// TODO... Готов byteJSON!
		}
	}
}
