package logg

import (
	"context"
	"log/slog"
)

type LoggKafka struct {
	store    chan slog.Record
	capQueue int
}

func NewLoggKafka(ctx context.Context, cap int) *LoggKafka {
	return &LoggKafka{
		store:    make(chan slog.Record, cap),
		capQueue: cap,
	}
}

func (l *LoggKafka) sendbuff(ctx context.Context, r slog.Record) {
	select {
	case <-ctx.Done():
		// Проверка наличия данных в канале.
		if len(l.store) > 0 {
			// Доработать тут . Сервер падает если
			// Если данные есть отправляем их в брокер.
		}
		return
	case l.store <- r:
	default:
		// Бдуем ли мы что нить откидывать лог
	}
}

func (l *LoggKafka) Send(ctx context.Context, r slog.Record) {
	// Отправка в буфер канала
	go l.sendbuff(ctx, r)

	return
}
