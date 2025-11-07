package server

import (
	"mealmate/cmd/config"
	"mealmate/internal/logg"
	"mealmate/internal/store"
)

func Start(
	config config.Config,
	appLog logg.Logger,
	infraLog logg.Logger,
	bsnessLog logg.Logger,
	storeDB store.DataBase) {

	// Router
	router := NewRouter()

	// Server
	NewServer(config, appLog).Run()

}
