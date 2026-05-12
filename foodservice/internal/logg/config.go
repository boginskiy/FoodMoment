package logg

type Config struct {
	Path         string // Путь к файлу (пусто = stdout)
	Level        string // debug, info, warn, error
	AddSource    bool   // Инфа по вызову логирования (какой файл, строка, функция)
	KafkaEnabled bool   // Использование передачи логов в kafka
}

var (
	// Config for main logger.
	MainLogCfg = &Config{
		Path:      "main.log",
		Level:     "INFO",
		AddSource: true,
	}

	// Config for metrics logger.
	MetricsLogCfg = &Config{
		Path:         "metrics.log",
		Level:        "INFO",
		AddSource:    false,
		KafkaEnabled: true,
	}

	// Config for kafka logger.
	KafkaLogCfg = &Config{
		Path:      "kafka.log",
		Level:     "INFO",
		AddSource: false,
	}
)
