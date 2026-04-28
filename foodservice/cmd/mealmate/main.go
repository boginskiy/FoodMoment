package main

import (
	"mealmate/cmd/config"
	"mealmate/cmd/server"
	"mealmate/internal/logg"
	"mealmate/internal/store"
)

var LEVEL = "INFO"

func main() {
	// Base Logger
	appLog := logg.NewLogg("app.log", LEVEL)

	// Config
	config := config.NewArgsENV(appLog)

	// Extra loggers
	bsnessLog := logg.NewLogg(config.GetBsnessLog(), LEVEL)
	infraLog := logg.NewLogg(config.GetInfraLog(), LEVEL)

	// Database
	storeDB := store.NewStoreDB(config, infraLog)

	// Defer
	defer bsnessLog.Close()
	defer infraLog.Close()
	defer appLog.Close()

	server.Start(config, appLog, infraLog, bsnessLog, storeDB)
}
