package config

import "mealmate/internal/logg"

type ArgsENV struct {
	RunAddress string `env:"RUN_ADDRESS"`
	NameLog    string `env:"NAME_LOG"`
	// NameWarnLog string `env:"NAME_WARN_LOG"`
	LevelLog string `env:"LEVEL_LOG"`
	Logger   logg.Logger
}

func NewArgsENV(logger logg.Logger) *ArgsENV {
	return &ArgsENV{Logger: logger}
}

func (a *ArgsENV) GetRunAddress() string {
	return ":8080"
}

func (a *ArgsENV) GetBsnessLog() string {
	return "bsness.log"
}

func (a *ArgsENV) GetInfraLog() string {
	return "infra.log"
}
