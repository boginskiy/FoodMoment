package routes

import (
	"mealmate/internal/handlers"

	"github.com/go-chi/chi"
)

type AuthRoutes struct {
	AuthHandler handlers.AuthHandler
}

func NewAuthRoutes(authHandler handlers.AuthHandler) *AuthRoutes {
	return &AuthRoutes{AuthHandler: authHandler}
}

func (a *AuthRoutes) RegisterRoutes(route chi.Router) {
	route.Post("/reset-password", a.AuthHandler.ResetPass) // Восстановление пароля
	route.Post("/register", a.AuthHandler.Register)        // Регистрация
	route.Post("/logout", a.AuthHandler.Logout)            // Выход
	route.Post("/login", a.AuthHandler.Login)              // Вход
}
