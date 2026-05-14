package app

import (
	"context"
	"net"

	"github.com/boginskiy/FoodMoment/foodservice/cmd/config"
	"github.com/boginskiy/FoodMoment/foodservice/internal/kafkalogg"
	"github.com/boginskiy/FoodMoment/foodservice/internal/logg"
	"github.com/boginskiy/FoodMoment/foodservice/internal/server"
	"github.com/boginskiy/FoodMoment/foodservice/pkg"
)

type App struct {
	Config      config.Config    // Config
	Logger      logg.Logger      // Logger main log
	KLogger     logg.Logger      // KLogger which write log about working Kafka
	MLogger     logg.Logger      // MLogger which write log about metrics
	KafkaSender kafkalogg.Sender // KafkaSender
	Server      server.Server    //

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
	// + Handler
	return a.Server.Run(ctx)
}

func (a *App) Close() error {
	a.KLogger.Close()
	a.KafkaSender.Close()
	a.MLogger.Close()
	a.MLogger.Close()
	a.Logger.Close()
	return nil
}

func (a *App) InitModules(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,      // Config главный
		a.initLogger,      // Logger главный
		a.initKLogger,     // Logger, который контролирует работу брокера
		a.initKafkaSender, // Sender отправщик данных в kafka
		a.initMLogger,     // Logger для сборка статистики и отправка в брокер
		a.initServer,      // Server

		// Репозитории ?
		// Сервисы ?
		// Хендлеры ?
		// Роутеры
		// Сервер in progress...
	}

	for _, init := range inits {
		err := init(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initServer(ctx context.Context) error {
	// Config
	addr := net.JoinHostPort(
		a.Config.GetString("server_host", "localhost"),
		a.Config.GetString("server_port", "8080"),
	)

	cfg := server.NewConfig(
		a.Config.GetString(addr, server.HTTPCfg.Addr),
		a.Config.GetDuration("read_timeout", server.HTTPCfg.ReadTimeout),
		a.Config.GetDuration("write_timeout", server.HTTPCfg.WriteTimeout),
		a.Config.GetDuration("idle_timeout", server.HTTPCfg.IdleTimeout),
	)

	s, err := server.NewHTTPServer(ctx, cfg, a.Logger)
	if err != nil {
		return err
	}
	a.Server = s
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

func (a *App) initKLogger(ctx context.Context) (err error) {
	// Config
	cfg := logg.Config{
		Path:      a.Config.GetString("path_kafka_log", logg.KafkaLogCfg.Path),
		Level:     a.Config.GetString("level_kafka_log", logg.KafkaLogCfg.Level),
		AddSource: a.Config.GetBool("addsource_kafka_log", logg.KafkaLogCfg.AddSource),
	}

	a.KLogger, err = logg.NewLogg(cfg, logg.JSONHandler)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initKafkaSender(ctx context.Context) (err error) {
	// Kafka config
	kcfg := kafkalogg.Config{
		Servers:     a.Config.GetString("servers_kafka_log", kafkalogg.Cfg.Servers),
		ClientID:    a.Config.GetString("clientID_kafka_log", kafkalogg.Cfg.ClientID),
		Acks:        a.Config.GetString("acks_kafka_log", kafkalogg.Cfg.Acks),
		Topic:       a.Config.GetString("topic_kafka_log", kafkalogg.Cfg.Topic), // Нет в config.json
		ChannelSize: a.Config.GetInt("channelsize_kafka_log", kafkalogg.Cfg.ChannelSize),
		MaxMessages: a.Config.GetInt("maxmessages_kafka_log", kafkalogg.Cfg.MaxMessages),
	}
	// Sender in Kafka.
	a.KafkaSender, err = kafkalogg.NewKafkaSend(ctx, kcfg, a.KLogger)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initMLogger(ctx context.Context) (err error) {
	// Config
	cfg := logg.Config{
		Path:         a.Config.GetString("path_metrics_log", logg.MetricsLogCfg.Path),
		Level:        a.Config.GetString("level_metrics_log", logg.MetricsLogCfg.Level),
		AddSource:    a.Config.GetBool("addsource_metrics_log", logg.MetricsLogCfg.AddSource),
		KafkaEnabled: a.Config.GetBool("enabled_kafkalog", logg.MetricsLogCfg.KafkaEnabled),
	}

	a.MLogger, err = kafkalogg.NewKafkaLogger(cfg, logg.JSONHandler, a.KafkaSender)
	if err != nil {
		return err
	}
	return nil
}
