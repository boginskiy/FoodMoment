package logg

import (
	"context"
	"log/slog"
)

type ProducLogger interface {
	Send(context.Context, slog.Record)
	Close() error
}

type Logger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
	With(args ...any) *Logg
	Close() error
}
