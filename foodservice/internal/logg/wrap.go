package logg

import (
	"io"
	"log/slog"
)

// HandlerFactory — тип функции, создающей хендлер
type HandlerFactory func(w io.Writer, opts *slog.HandlerOptions) slog.Handler

func JSONHandler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	return slog.NewJSONHandler(w, opts)
}

func TextHandler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	return slog.NewTextHandler(w, opts)
}
