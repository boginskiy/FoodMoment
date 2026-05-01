package config

import (
	"context"
	"flag"

	"github.com/caarlos0/env"
)

type ServConfig struct {
	Host string `env:"Food_HOST"`
	Port string `env:"Food_PORT"`
	// ReadTimeout  time.Duration
	// WriteTimeout time.Duration
}

func NewServConfig(ctx context.Context) (*ServConfig, error) {
	tmpCfg := &ServConfig{}

	// Parsing env. Stage 1. ENV.
	err := env.Parse(&tmpCfg)
	if err != nil {
		return nil, err
	}

	// Parsing flag. Stage 2. CLI.
	tmpCfg.parseFlags(ctx)

	// Load JSON. Stage 3. File JSON.

	return tmpCfg, nil
}

func (s *ServConfig) parseFlags(ctx context.Context) {
	// Для пустых атрибутов пробуем спарсить значения флагов с CLI.
	switch {
	case s.Host == "":
		flag.StringVar(&s.Host, "h", "", "...")
		fallthrough

	case s.Port == "":
		flag.StringVar(&s.Port, "p", "", "...")
	}
	flag.Parse()

	// А будет ли такое работать ??
}
