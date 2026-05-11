package app

import (
	"context"

	"github.com/boginskiy/FoodMoment/foodservice/cmd/config"
	"github.com/boginskiy/FoodMoment/foodservice/internal/logg"
	"github.com/boginskiy/FoodMoment/foodservice/pkg"
)

type App struct {
	Config      config.Config
	Logger      logg.Logger
	KafkaLogger logg.Logger
	SLogger     logg.Logger
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
		a.initKafkaLogger,
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

func (a *App) initLogger(ctx context.Context) (err error) {
	// Config
	cfg := logg.Config{
		Path:      a.Config.GetString("path_main_log", logg.MainLogCfg.Path),
		Level:     a.Config.GetString("level_main_log", logg.MainLogCfg.Level),
		AddSource: a.Config.GetBool("addsource_main_log", logg.MainLogCfg.AddSource),
	}

	a.Logger, err = logg.NewLogg(cfg, logg.JSONHandler)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initKafkaLogger(ctx context.Context) (err error) {
	// Config
	cfg := logg.Config{
		Path:      a.Config.GetString("path_kafka_log", logg.KafkaLogCfg.Path),
		Level:     a.Config.GetString("level_kafka_log", logg.KafkaLogCfg.Level),
		AddSource: a.Config.GetBool("addsource_kafka_log", logg.KafkaLogCfg.AddSource),
	}

	a.KafkaLogger, err = logg.NewLogg(cfg, logg.JSONHandler)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initSLogger(ctx context.Context) error {
	// Config
	cfg := logg.Config{
		Path:         a.Config.GetString("path_metrics_log", logg.MetricsLogCfg.Path),
		Level:        a.Config.GetString("level_metrics_log", logg.MetricsLogCfg.Level),
		AddSource:    a.Config.GetBool("addsource_metrics_log", logg.MetricsLogCfg.AddSource),
		KafkaEnabled: a.Config.GetBool("enabled_kafkalog", logg.MetricsLogCfg.KafkaEnabled),
	}

	// Sender
	sender, err := a.newKafkaSender(ctx, a.KafkaLogger)
	if err != nil {
		return err
	}

	a.SLogger, err = logg.NewKafkaLogger(cfg, logg.JSONHandler, sender)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) newKafkaSender(ctx context.Context, log logg.Logger) (*logg.KafkaSend, error) {
	// Kafka config
	kcfg := logg.KafkaConfig{
		Servers:     a.Config.GetString("servers_kafka_log", logg.KafkaCfg.Servers),
		ClientID:    a.Config.GetString("clientID_kafka_log", logg.KafkaCfg.ClientID),
		Acks:        a.Config.GetString("acks_kafka_log", logg.KafkaCfg.Acks),
		ChannelSize: a.Config.GetInt("channelsize_kafka_log", logg.KafkaCfg.ChannelSize),
		MaxMessages: a.Config.GetInt("maxmessages_kafka_log", logg.KafkaCfg.MaxMessages),
	}
	// Sender in Kafka.
	return logg.NewKafkaSend(ctx, kcfg, log)
}

// TODO.
// Логгер, далее методы, кафку продумать, откидываем ошибки уровня error
// Далее сервер делать
// Роутер
// и т.п.
