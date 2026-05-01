package config

import (
	"context"
	"flag"
	"os"

	"github.com/boginskiy/FoodMoment/foodservice/pkg"
)

type Config struct {
	Server   ServerConfig
	DataBase DataBaseConfig
	Auth     AuthConfig
	Redis    RedisConfig

	loader pkg.LoaderConfig
}

func NewConfig(ctx context.Context, loader pkg.LoaderConfig) (*Config, error) {
	tmpCfg := &Config{}

	// LoaderConfig.
	loader, _ := pkg.NewLoadConfig(ctx, tmpCfg.getPathToConfigJSON())

	err := tmpCfg.InitConfig(ctx)
	if err != nil {
		return nil, err
	}

	return tmpCfg, nil
}

func (c *Config) InitConfig(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		c.InitServerConfig,
	}

	for _, init := range inits {
		err := init(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Config) InitServerConfig(ctx context.Context) error {
	var err error
	c.Server, err = NewServConfig(ctx)
	if err != nil {
		return err
	}
	return nil
}
