package kafkalogg

type Config struct {
	Servers     string // Минимум 2-3 брокера для отказоустойчивости
	ClientID    string // Логическое имя продюсера, которое отправляется в каждом запросе к брокерам
	Acks        string // Компромисс. Только лидер партиции подтверждает получение данных. Реплики могут не успеть получить данные
	Topic       string //
	ChannelSize int    // Оптимизация производительности
	MaxMessages int    // Оптимизация производительности

}

var (
	// Kafka Server Settings.
	Cfg = &Config{
		Servers:     "kafka1:9092,kafka2:9092,kafka3:9092",
		ClientID:    "KafkaLogger",
		Topic:       "logs",
		Acks:        "1",
		ChannelSize: 1000,
		MaxMessages: 1000,
	}

	// Время обработки Response Message после остановки сервиса.
	workingTimeAfterDone = 3
)
