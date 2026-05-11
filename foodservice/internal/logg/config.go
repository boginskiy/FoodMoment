package logg

type Config struct {
	Path         string // Путь к файлу (пусто = stdout)
	Level        string // debug, info, warn, error
	AddSource    bool   // Инфа по вызову логирования (какой файл, строка, функция)
	KafkaEnabled bool   // Использование передачи логов в kafka
}

type KafkaConfig struct {
	Servers     string // Минимум 2-3 брокера для отказоустойчивости
	ClientID    string // Логическое имя продюсера, которое отправляется в каждом запросе к брокерам
	Acks        string // Компромисс. Только лидер партиции подтверждает получение данных. Реплики могут не успеть получить данные
	ChannelSize int    // Оптимизация производительности
	MaxMessages int    // Оптимизация производительности

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

	// Kafka Server Settings.
	KafkaCfg = &KafkaConfig{
		Servers:     "kafka1:9092,kafka2:9092,kafka3:9092",
		ClientID:    "KafkaLogger",
		Acks:        "1",
		ChannelSize: 1000,
		MaxMessages: 1000,
	}

	// Время обработки Response Message после остановки сервиса.
	workingTimeAfterDone = 3
)
