package config

import (
	"context"
	"flag"
)

type CliProvider struct {
	parsed bool
}

func NewCliProvider(ctx context.Context) *CliProvider {
	return &CliProvider{parsed: false}
}

func (f *CliProvider) Load(key string) (any, bool) {
	if !f.parsed {
		flag.Parse()
		f.parsed = true
	}
	if f := flag.Lookup(key); f != nil {
		return f.Value.String(), true
	}
	return nil, false
}
