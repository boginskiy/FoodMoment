package kafkalogg

import (
	"context"
	"log/slog"
)

type Preparer interface {
	SerializeToJSON(slog.Record) ([]byte, error)
}

type Writer interface {
	Send(topic string, msg []byte)
	Close() error
}

type Sender interface {
	Send(context.Context, slog.Record)
	Close() error
}
