package logg

import (
	"io"
	"log/slog"
	"os"
)

type Logg struct {
	*slog.Logger
	closer io.Closer
}

func NewLogg(cfg Config, factory HandlerFactory) (*Logg, error) {
	var writer io.Writer = os.Stderr
	var closer io.Closer

	// Writer / Closer
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

	// Handler
	handler := factory(writer, opt)

	return &Logg{
		Logger: slog.New(handler),
		closer: closer,
	}, nil
}

func (l *Logg) Close() error {
	if l.closer != nil {
		return l.closer.Close()
	}
	return nil
}

func (l *Logg) With(args ...any) *Logg {
	return &Logg{
		Logger: l.Logger.With(args...),
		closer: l.closer,
	}
}

func (l *Logg) Info(msg string, args ...any) {
	l.Logger.Info(msg, args...)
}

func (l *Logg) Error(msg string, args ...any) {
	l.Logger.Error(msg, args...)
}
