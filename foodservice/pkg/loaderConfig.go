package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
)

type LoaderConfig interface {
	GetConfig(name string) (any, error)
}

type LoadConfig struct {
	Store  map[string]any
	Getter GetterValue
}

var (
	varENVPathCfgFile = "PATH_CONFIG"
	varCLIPathCfgFile = "path"
)

func NewLoadConfig(ctx context.Context, getter GetterValue) (*LoadConfig, error) {
	tmp := &LoadConfig{
		Store:  map[string]any{},
		Getter: getter,
	}

	// TODO...
	// Остановка тут. Надо понимать что делать с передчаей парамтеров - varENVPathCfgFile, varCLIPathCfgFile
	// Создать еще структура для чтнеия // ReaderFile
	// Схема json, строгая или оставляем на интерфейсах ?

	path := tmp.Getter.GetStringValue(varENVPathCfgFile, varCLIPathCfgFile)

	bytes, err := os.ReadFile(path)
	if err != nil {
		return tmp, err
	}

	err = json.Unmarshal(bytes, &tmp.Store)
	if err != nil {
		return tmp, err
	}

	return tmp, nil
}

func (l *LoadConfig) GetConfig(name string) (any, error) {
	val, ok := l.Store[name]
	if !ok {
		return nil, fmt.Errorf("no data available in config file")
	}
	return val, nil
}
