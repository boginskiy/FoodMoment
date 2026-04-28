package config

type Config interface {
	GetRunAddress() string
	GetBsnessLog() string
	GetInfraLog() string
}
