package config

import "context"

type Config struct {
}

func NewConfig(ctx context.Context) (*Config, error) {
	tmpCfg := &Config{}
	return tmpCfg, nil
}

// TODO...
// Сущности? Какой приоритет по пересекающимся параметрам?
//
