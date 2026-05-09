package app

import (
	"context"
	"log/slog"

	"github.com/boginskiy/FoodMoment/foodservice/cmd/config"
	"github.com/boginskiy/FoodMoment/foodservice/internal/logg"
	"github.com/boginskiy/FoodMoment/foodservice/pkg"
)

type App struct {
	Config  config.Config
	Logger  logg.Logger
	SLogger *slog.Logger
}

func NewApp(ctx context.Context) (*App, error) {
	tmpApp := &App{}

	// Init modules.
	err := tmpApp.InitModules(ctx)
	if err != nil {
		return nil, err
	}
	return tmpApp, nil
}

func (a *App) Run(ctx context.Context) error {
	return nil
}

func (a *App) InitModules(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initLogger,
		a.initSLogger,
	}

	for _, init := range inits {
		err := init(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initLogger(ctx context.Context) error {
	config := logg.Config{
		Path:      a.Config.GetString("path_main_log", "main.log"),
		Level:     a.Config.GetString("level_main_log", "INFO"),
		AddSource: a.Config.GetBool("addsource_main_log", true),
	}
	var err error
	a.Logger, err = logg.NewLogg(config, logg.JSONHandler)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initConfig(ctx context.Context) error {
	// Create providers
	jsonProvider := config.NewJSONProvider(ctx, pkg.NewGetVar(ctx), pkg.NewReadFile(ctx))
	cliProvider := config.NewCliProvider(ctx)
	envProvider := config.NewEnvProvider(ctx)

	var err error
	a.Config, err = config.NewConf(ctx, envProvider, cliProvider, jsonProvider)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initSLogger(ctx context.Context) error {
	config := logg.Config{
		Path:         a.Config.GetString("path_infra_log", "infra.log"),
		Level:        a.Config.GetString("level_infra_log", "INFO"),
		AddSource:    a.Config.GetBool("addsource_infra_log", false),
		KafkaEnabled: a.Config.GetBool("kafka_enabled_log", true),
		KafkaLevel:   a.Config.GetString("kafka_level_log", "INFO"),
	}

	// Producer by Kafka.
	produceLogger := logg.NewProduceLogg(ctx, a.Logger, CapQueue)

	var err error
	a.SLogger, err = logg.NewSLogg(config, logg.JSONHandler, produceLogger)
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
