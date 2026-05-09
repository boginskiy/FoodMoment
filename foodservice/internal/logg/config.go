package logg

type Config struct {
	Path         string // Путь к файлу (пусто = stdout)
	Level        string // debug, info, warn, error
	AddSource    bool   // Инфа по вызову логирования (какой файл, строка, функция)
	KafkaEnabled bool   // Kafka
	KafkaLevel   string // Kafka
}
