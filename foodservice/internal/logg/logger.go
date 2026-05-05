package logg

import (
	"io"
	"log/slog"
	"os"
)

type OutWrite struct {
	io.Writer
}

func NewOutWrite(path string) (*OutWrite, error) {
	stdOut := os.Stdout
	if path != "" {
		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return nil, err
		}
		stdOut = file
	}
	return &OutWrite{stdOut}, nil
}

type LoggJSON struct {
	Handler *slog.JSONHandler
}

func NewLoggJSON(out io.Writer, options *slog.HandlerOptions) *LoggJSON {
	return &LoggJSON{
		Handler: slog.NewJSONHandler(out, options),
	}
}

type Logg struct {
	Handler slog.Handler
}

func NewLogg(handler slog.Handler) *Logg {
	return &Logg{
		Handler: handler,
	}
}
