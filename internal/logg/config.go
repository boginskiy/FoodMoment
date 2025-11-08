package logg

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var LEVEL = map[string]zapcore.Level{
	"INFO":    zap.InfoLevel,
	"WARNING": zap.WarnLevel,
	"ERROR":   zap.ErrorLevel,
	"FATAL":   zap.FatalLevel,
}

func Config(file, level string) zap.Config {
	// NewProductionConfig / NewDevelopmentConfig
	cfg := zap.NewProductionConfig()

	// Уровень логирования
	cfg.Level.SetLevel(LEVEL[level])

	// Перенаправление логов
	cfg.OutputPaths = []string{"stdout", file}

	cfg.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "time",                        // Ключевое поле времени ("time")
		MessageKey:     "msg",                         // Ключевое поле сообщения ("msg")
		LevelKey:       "lvl",                         // Ключевое поле уровня лога ("lvl")
		NameKey:        "",                            // Имя пакета игнорируется
		CallerKey:      "",                            // Название исходящего места ("caller")
		FunctionKey:    "",                            // Функция игнорируется
		StacktraceKey:  "",                            // Стэк-трейс, если возникает ошибка ("stacktrace")
		LineEnding:     zapcore.DefaultLineEnding,     // Перенос строки по умолчанию
		EncodeTime:     zapcore.RFC3339TimeEncoder,    // формат времени
		EncodeDuration: zapcore.StringDurationEncoder, // Длительность в секундах
		EncodeCaller:   zapcore.ShortCallerEncoder,    // Сокращённое представление caller
		EncodeLevel:    zapcore.CapitalLevelEncoder,
	}
	return cfg
}
