package config

import (
	"context"
	"strconv"
	"time"
)

type Conf struct {
	provider ValuePriority
}

func NewConf(ctx context.Context, envProvider, cliProvider, jsonProvider ValueProvider) (*Conf, error) {
	providers := []ValueProvider{
		envProvider,  // приоритет 1
		cliProvider,  // приоритет 2
		jsonProvider, // приоритет 3
	}

	return &Conf{
		provider: NewPriorityProvider(providers...),
	}, nil
}

func (c *Conf) GetString(key, defaultValue string) string {
	if val := c.provider.Load(key); val != nil {
		if v, ok := val.(string); ok {
			return v
		}
	}
	return defaultValue
}

func (c *Conf) GetInt(key string, defaultValue int) int {
	if val := c.provider.Load(key); val != nil {

		switch v := val.(type) {
		case int:
			return v
		case float64:
			return int(v)
		case string:
			if i, err := strconv.Atoi(v); err == nil {
				return i
			}
		}
	}
	return defaultValue
}

func (c *Conf) GetBool(key string, defaultValue bool) bool {
	if val := c.provider.Load(key); val != nil {
		if v, ok := val.(bool); ok {
			return v
		}
	}
	return defaultValue
}

func (c *Conf) GetDuration(key string, defaultValue time.Duration) time.Duration {
	if val := c.provider.Load(key); val != nil {
		if v, ok := val.(int); ok {
			return time.Duration(v) * time.Second
		}
	}
	return defaultValue
}
