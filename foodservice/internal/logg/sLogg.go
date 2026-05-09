package logg

import (
	"context"
	"io"
	"log/slog"
	"os"
)

func NewSLogg(cfg Config, factory HandlerFactory, producer ProducLogger) (*slog.Logger, error) {
	var writer io.Writer = os.Stderr
	var closer io.Closer

	// Writer
	if cfg.Path != "" {
		file, err := os.OpenFile(cfg.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return nil, err
		}
		writer = file
		closer = file
	}

	// Options
	opt := &slog.HandlerOptions{
		Level:     Level[cfg.Level],
		AddSource: cfg.AddSource,
	}

	// Базовый логгер. Без поддержки Kafka.
	handler := factory(writer, opt)
	finalHandler := handler

	// Расширенный логгер. С поддержкой Kafka.
	if producer != nil && cfg.KafkaEnabled {
		finalHandler = &KafkaHandler{
			Producer:    producer,
			Config:      cfg,
			baseHandler: handler,
			closer:      closer,
		}
	}

	return slog.New(finalHandler), nil
}

type KafkaHandler struct {
	Producer    ProducLogger
	Config      Config
	baseHandler slog.Handler
	closer      io.Closer
}

func (k *KafkaHandler) Handle(ctx context.Context, r slog.Record) error {
	// Отправка логов в Kafka с учетом заданного уровня.
	if k.Producer != nil && r.Level == Level[k.Config.KafkaLevel] {
		go func() {
			// TODO. Использование отдельного контекста?
			k.Producer.Send(ctx, r)

			// if err != nil {
			// 	// Логируем ошибку в основной логгер.
			// 	k.baseHandler.Handle(ctx, slog.NewRecord(
			// 		r.Time,
			// 		slog.LevelWarn,
			// 		"failed to send log to Kafka",
			// 		0,
			// 	))
			// }
		}()
	}

	// Передаем далее
	if k.baseHandler != nil {
		return k.baseHandler.Handle(ctx, r)
	}
	return nil
}

func (k *KafkaHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return k.baseHandler.Enabled(ctx, level)
}

func (k *KafkaHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &KafkaHandler{
		baseHandler: k.baseHandler.WithAttrs(attrs),
		Producer:    k.Producer,
		Config:      k.Config,
		closer:      k.closer,
	}
}

func (k *KafkaHandler) WithGroup(name string) slog.Handler {
	return &KafkaHandler{
		baseHandler: k.baseHandler.WithGroup(name),
		Producer:    k.Producer,
		Config:      k.Config,
		closer:      k.closer,
	}
}
