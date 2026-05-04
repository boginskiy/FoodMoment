package config

type ValueProvider interface {
	Load(key string) (any, bool)
}

type ValuePriority interface {
	Load(key string) any
}

type Config interface {
	GetString(key, defaultValue string) string
	GetInt(key string, defaultValue int) int
}
