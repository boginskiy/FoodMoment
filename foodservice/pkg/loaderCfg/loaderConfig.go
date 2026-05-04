package loaderCfg

import (
	"context"
	"fmt"
)

type ConfigLoader interface {
	GetConfig(name string) (any, error)
}

type LoadConfig struct {
	Store  map[string]any
	Getter FlagEnvGetter
	Reader FileReader
}

func NewLoadConfig(ctx context.Context) (*LoadConfig, error) {
	tmp := &LoadConfig{
		Store:  map[string]any{},
		Getter: NewGetVar(ctx),
		Reader: NewReadFile(ctx),
	}

	// Const variables.
	path := tmp.choosePath(nameOfVarPathCfgFileCLI, nameOfVarPathCfgFileENV, defaultPathCfgFile)

	// Read config.json.
	err := tmp.Reader.Deserialization(path, &tmp.Store)
	if err != nil {
		return nil, err
	}

	return tmp, nil
}

func (l *LoadConfig) choosePath(flagKey, envKey, defaultVal string) string {
	if path := l.Getter.GetStringValueFromENV(envKey); path != "" {
		return path
	}
	if path := l.Getter.GetStringValueFromCLI(flagKey); path != "" {
		return path
	}
	return defaultVal
}

func (l *LoadConfig) GetConfig(name string) (any, error) {
	val, ok := l.Store[name]
	if !ok {
		return nil, fmt.Errorf("no data available in config file")
	}
	return val, nil
}
