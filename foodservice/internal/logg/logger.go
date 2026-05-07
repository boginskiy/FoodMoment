package logg

import (
	"context"
	"io"
	"log/slog"
	"os"
)

type Logg struct {
	superlog slog.Handler
	producer LoggerKafka
	level    string
}

func NewLogg(ctx context.Context, path, levelFile, levelKafka string, factory HandlerFactory) (*slog.Logger, error) {
	tmp := &Logg{level: levelKafka}

	// Writer
	writer, err := tmp.createWriter(path)
	if err != nil {
		return nil, err
	}
	// Options
	options := &slog.HandlerOptions{
		Level: Level[levelFile],
	}

	tmp.superlog = factory(writer, options)

	return slog.New(tmp), nil
}

func (l *Logg) Handle(ctx context.Context, r slog.Record) error {
	// Отправка логов в Kafka с учетом заданного уровня.
	if r.Level == Level[l.level] {
		l.producer.Send(r)
	}

	if l.superlog != nil {
		return l.superlog.Handle(ctx, r)
	}
	return nil
}

func (l *Logg) Enabled(ctx context.Context, level slog.Level) bool {
	return l.superlog.Enabled(ctx, level)
}

func (l *Logg) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Logg{
		superlog: l.superlog.WithAttrs(attrs),
		producer: l.producer,
		level:    l.level,
	}
}

func (l *Logg) WithGroup(name string) slog.Handler {
	return &Logg{
		superlog: l.superlog.WithGroup(name),
		producer: l.producer,
		level:    l.level,
	}
}

func (l *Logg) createWriter(path string) (io.Writer, error) {
	writer := os.Stdout
	if path != "" {
		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return nil, err
		}
		writer = file
	}
	return writer, nil
}
