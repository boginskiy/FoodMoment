package pkg

import (
	"context"
	"flag"
	"os"
)

type GetterValue interface {
	GetStringValue(envKey, flagKey string) string
}

type FlagEnvGetter struct {
	parsed bool
}

func NewGetVar(ctx context.Context) *FlagEnvGetter {
	return &FlagEnvGetter{}
}

func (g *FlagEnvGetter) GetStringValue(envKey, flagKey string) string {
	// Пробуем ENV.
	if val := os.Getenv(envKey); val != "" {
		return val
	}

	// Пробуем CLI. Парсим.
	if !g.parsed {
		flag.Parse()
		g.parsed = true
	}

	if fl := flag.Lookup(flagKey); fl != nil {
		return fl.Value.String()
	}

	return ""
}
