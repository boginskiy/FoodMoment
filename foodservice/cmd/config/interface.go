package config

type ValueProvider interface {
	Load(key string) (any, bool)
}

type ValuePriority interface {
	Load(key string) any
}

type Config interface {
	GetString(key, defaultValue string) string
	GetBool(key string, defaultValue bool) bool
	GetInt(key string, defaultValue int) int
}
