package pkg

import (
	"context"
	"flag"
	"os"
)

type FlagEnvGetter interface {
	GetValueFromCLI(flagKey string) string
	GetValueFromENV(envKey string) string
}

type FlagEnvGet struct {
	parsed bool
}

func NewGetVar(ctx context.Context) *FlagEnvGet {
	return &FlagEnvGet{}
}

// GetValueFromCLI парсинг CLI.
func (g *FlagEnvGet) GetValueFromCLI(flagKey string) string {
	if !g.parsed {
		flag.Parse()
		g.parsed = true
	}

	if fl := flag.Lookup(flagKey); fl != nil {
		return fl.Value.String()
	}
	return ""
}

// GetValueFromENV парсинг ENV.
func (g *FlagEnvGet) GetValueFromENV(envKey string) string {
	if val := os.Getenv(envKey); val != "" {
		return val
	}
	return ""
}
