package store

import (
	"mealmate/cmd/config"
	"mealmate/internal/logg"
)

type StoreDB struct {
	Cfg  config.Config
	Logg logg.Logger
}

func NewStoreDB(config config.Config, logger logg.Logger) *StoreDB {
	return &StoreDB{Cfg: config, Logg: logger}
}
