package app

import (
	"context"
	"log/slog"

	"github.com/boginskiy/FoodMoment/foodservice/cmd/config"
	"github.com/boginskiy/FoodMoment/foodservice/internal/logg"
	"github.com/boginskiy/FoodMoment/foodservice/pkg"
)

type App struct {
	Config   config.Config
	MainLog  *slog.Logger
	InfraLog *slog.Logger
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

func (a *App) initConfig(ctx context.Context) (*config.Conf, error) {
	// Create providers
	jsonProvider := config.NewJSONProvider(ctx, pkg.NewGetVar(ctx), pkg.NewReadFile(ctx))
	cliProvider := config.NewCliProvider(ctx)
	envProvider := config.NewEnvProvider(ctx)

	return config.NewConf(ctx, envProvider, cliProvider, jsonProvider)
}

func (a *App) initModules(ctx context.Context) error {
	// Config
	cfg, err := a.initConfig(ctx)
	if err != nil {
		return err
	}

	inits := []func(context.Context, config.Config) error{
		a.initMainLogger,
		a.initInfraLogger,
	}

	for _, init := range inits {
		err := init(ctx, cfg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initMainLogger(ctx context.Context, cfg config.Config) error {
	path := cfg.GetString("path_log", config.MainLog)
	levelF := cfg.GetString("level_file_log", config.LevelWarn)
	levelK := cfg.GetString("level_kafka_log", config.LevelInfo)

	var err error

	a.MainLog, err = logg.NewLogg(ctx, path, levelF, levelK, logg.JSONHandler)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initInfraLogger(ctx context.Context, cfg config.Config) error {
	path := ""
	levelF := "DEBUG"
	levelK := "INFO"

	var err error

	a.InfraLog, err = logg.NewLogg(ctx, path, levelF, levelK, logg.TextHandler)
	if err != nil {
		return err
	}
	return nil
}

// TODO.
// Логгер, далее методы, кафку продумать, откидываем ошибки уровня error

// Далее сервер делать
// Роутер
// и т.п.
