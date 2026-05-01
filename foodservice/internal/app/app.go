package app

import (
	"context"

	"github.com/boginskiy/FoodMoment/foodservice/cmd/config"
	"github.com/boginskiy/FoodMoment/foodservice/pkg"
)

type App struct {
	cfg *config.Config
}

func NewApp(ctx context.Context) (*App, error) {
	tmpApp := &App{}

	// Init modules.
	err := tmpApp.initModules(ctx)
	if err != nil {
		return nil, err
	}
	return tmpApp, nil
}

func (a *App) initModules(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
	}

	for _, init := range inits {
		err := init(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initConfig(ctx context.Context) error {
	var err error

	// GetterStringVar
	getterStringVar := pkg.NewGetVar(ctx)

	// LoaderConfig
	loaderConfig, _ := pkg.NewLoadConfig(ctx, getterStringVar)

	a.cfg, err = config.NewConfig(ctx)
	if err != nil {
		return err
	}
	return nil
}
