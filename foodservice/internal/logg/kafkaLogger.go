package logg

import (
	"context"
	"io"
	"log/slog"
	"os"
)

type KafkaLogg struct {
	slog   *slog.Logger
	closer io.Closer
}

func NewKafkaLogger(cfg Config, factory HandlerFactory, sender KafkaSender) (*KafkaLogg, error) {
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

	handler := factory(writer, opt)        // Базовый логгер. Без поддержки Kafka.
	if sender != nil && cfg.KafkaEnabled { // Расширенный логгер. С поддержкой Kafka.
		handler = &KafkaHandler{
			Sender:      sender,
			baseHandler: handler,
		}
	}

	return &KafkaLogg{slog: slog.New(handler), closer: closer}, nil
}

func (k *KafkaLogg) Close() error {
	if k.closer != nil {
		return k.closer.Close()
	}
	return nil
}

func (k *KafkaLogg) With(args ...any) Logger {
	return &KafkaLogg{
		slog:   k.slog.With(args...),
		closer: k.closer,
	}
}

func (k *KafkaLogg) Info(msg string, args ...any) {
	k.slog.Info(msg, args...)
}

func (k *KafkaLogg) Error(msg string, args ...any) {
	k.slog.Error(msg, args...)
}

// KafkaHandler
type KafkaHandler struct {
	Sender      KafkaSender
	baseHandler slog.Handler
}

func (k *KafkaHandler) Handle(ctx context.Context, r slog.Record) error {
	// Отправка логов в Kafka.
	if k.Sender != nil {
		go func() {
			// TODO. Использование отдельного контекста?
			k.Sender.Send(ctx, r)
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
		Sender:      k.Sender,
	}
}

func (k *KafkaHandler) WithGroup(name string) slog.Handler {
	return &KafkaHandler{
		baseHandler: k.baseHandler.WithGroup(name),
		Sender:      k.Sender,
	}
}
