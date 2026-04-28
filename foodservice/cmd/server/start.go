package server

import (
	"mealmate/cmd/config"
	"mealmate/internal/auth"
	"mealmate/internal/handlers"
	"mealmate/internal/logg"
	"mealmate/internal/middleware"
	"mealmate/internal/routes"
	"mealmate/internal/store"
)

func Start(
	config config.Config,
	appLog logg.Logger,
	infraLog logg.Logger,
	bsnessLog logg.Logger,
	storeDB store.DataBase) {

	// Auth
	auth := auth.NewAuth(config, appLog)

	// Middleware
	mdlwere := middleware.NewMdlware(config, appLog, auth)

	// Handlers
	authHandler := handlers.NewAuthHandler(config, appLog)

	// Routers
	authRoutes := routes.NewAuthRoutes(authHandler)
	router := NewRouter(authRoutes)

	// Server
	NewServer(config, appLog).Run(router, mdlwere)

}
