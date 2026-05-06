package logg

import (
	"io"
	"log/slog"
	"os"
)

type Logg struct {
	*slog.Logger
}

// HandlerFactory — тип функции, создающей хендлер
type HandlerFactory func(w io.Writer, opts *slog.HandlerOptions) slog.Handler

func NewLogg(path, lv string, factory HandlerFactory) (*Logg, error) {
	tmp := &Logg{}

	// Writer
	writer, err := tmp.createWriter(path)
	if err != nil {
		return nil, err
	}
	// Options
	options := &slog.HandlerOptions{
		Level: Level[lv],
	}

	tmp.Logger = slog.New(factory(writer, options))
	return tmp, nil
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

// TODO...далее...
func (l *Logg) RaiseInfo() {

}
