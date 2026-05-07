package logg

import "log/slog"

type LoggerKafka interface {
	Send(slog.Record)
}
