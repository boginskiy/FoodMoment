package logg

import (
	"context"
	"log/slog"
)

type Preparer interface {
	SerializeToJSON(slog.Record) ([]byte, error)
}

type KafkaWriter interface {
	Send(topic string, msg []byte)
	Close() error
}

type KafkaSender interface {
	Send(context.Context, slog.Record)
	Close() error
}

type Logger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
	With(args ...any) Logger
	Close() error
}
