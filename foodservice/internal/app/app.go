package app

import (
	"context"
	"log/slog"

	"github.com/boginskiy/FoodMoment/foodservice/cmd/config"
	"github.com/boginskiy/FoodMoment/foodservice/internal/logg"
	"github.com/boginskiy/FoodMoment/foodservice/pkg"
)

type App struct {
	Config config.Config
	Logger logg.Logger
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

func (a *App) Run(ctx context.Context) error {
	return nil
}

func (a *App) initModules(ctx context.Context) error {
	// Config
	cfg, err := a.initConfig(ctx)
	if err != nil {
		return err
	}

	inits := []func(context.Context, config.Config) error{
		a.initLogger,
	}

	for _, init := range inits {
		err := init(ctx, cfg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initConfig(ctx context.Context) (*config.Conf, error) {
	// Create providers
	jsonProvider := config.NewJSONProvider(ctx, pkg.NewGetVar(ctx), pkg.NewReadFile(ctx))
	cliProvider := config.NewCliProvider(ctx)
	envProvider := config.NewEnvProvider(ctx)

	return config.NewConf(ctx, envProvider, cliProvider, jsonProvider)
}

func (a *App) initLogger(ctx context.Context, cfg config.Config) error {
	options := &slog.HandlerOptions{
		Level: logg.Level[cfg.GetString("level_log", "INFO")],
	}

	outWrite, err := logg.NewOutWrite(cfg.GetString("path_log", "main.log"))
	if err ... // TODO ...

	return nil
}
