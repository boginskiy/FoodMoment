package app

import (
	"context"

	"github.com/boginskiy/FoodMoment/foodservice/cmd/config"
	"github.com/boginskiy/FoodMoment/foodservice/internal/kafkalogg"
	"github.com/boginskiy/FoodMoment/foodservice/internal/logg"
	"github.com/boginskiy/FoodMoment/foodservice/pkg"
)

type App struct {
	Config      config.Config
	Logger      logg.Logger
	KafkaLogger logg.Logger
	KafkaSender kafkalogg.Sender
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

func (a *App) Close() error {
	a.KafkaLogger.Close()
	a.KafkaSender.Close()
	a.SLogger.Close()
	a.SLogger.Close()
	a.Logger.Close()
	return nil
}

func (a *App) InitModules(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,      // Config главный
		a.initLogger,      // Logger главный
		a.initKafkaLogger, // Logger, который контролирует работу брокера
		a.initKafkaSender, // Sender отправщик данных в kafka
		a.initSLogger,     // Logger для сборка статистики и отправка в брокер
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
	a.KafkaSender, err = kafkalogg.NewKafkaSend(ctx, kcfg, a.KafkaLogger)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initSLogger(ctx context.Context) (err error) {
	// Config
	cfg := logg.Config{
		Path:         a.Config.GetString("path_metrics_log", logg.MetricsLogCfg.Path),
		Level:        a.Config.GetString("level_metrics_log", logg.MetricsLogCfg.Level),
		AddSource:    a.Config.GetBool("addsource_metrics_log", logg.MetricsLogCfg.AddSource),
		KafkaEnabled: a.Config.GetBool("enabled_kafkalog", logg.MetricsLogCfg.KafkaEnabled),
	}

	a.SLogger, err = kafkalogg.NewKafkaLogger(cfg, logg.JSONHandler, a.KafkaSender)
	if err != nil {
		return err
	}
	return nil
}

// TODO.
// Сделать минимальный функционал в этом сервисе и делать аутентификацию пользователя.
// Далеее сделать consumer для приемки метрик-логов черещ брокера

// TODO...
// Далее сервер делать
// Роутер
// и т.п.
