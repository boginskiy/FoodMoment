package config

import (
	"context"
	"os"
)

type EnvProvider struct {
}

func NewEnvProvider(ctx context.Context) *EnvProvider {
	return &EnvProvider{}
}

func (e *EnvProvider) Load(key string) (any, bool) {
	return os.LookupEnv(key)
}
