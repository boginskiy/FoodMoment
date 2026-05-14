package server

import "time"

type Config struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func NewConfig(addr string, readTm, writeTm, IdleTm time.Duration) Config {
	return Config{
		Addr:         addr,
		ReadTimeout:  readTm,
		WriteTimeout: writeTm,
		IdleTimeout:  IdleTm,
	}
}

var (
	HTTPCfg = &Config{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	defaultShutdownTimeout = 5 * time.Second
)
