package server

import (
	"mealmate/cmd/config"
	"mealmate/internal/handlers"
	"mealmate/internal/logg"
	"mealmate/internal/routes"
	"mealmate/internal/store"
)

func Start(
	config config.Config,
	appLog logg.Logger,
	infraLog logg.Logger,
	bsnessLog logg.Logger,
	storeDB store.DataBase) {

	// Handlers
	authHandler := handlers.NewAuthHandler(config, appLog)

	// Routers
	authRoutes := routes.NewAuthRoutes(authHandler)
	router := NewRouter(authRoutes)

	// Server
	NewServer(config, appLog).Run(router)

}
