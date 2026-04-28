package server

import (
	mv "mealmate/internal/middleware"
	"mealmate/internal/routes"

	"github.com/go-chi/chi"
)

type Route struct {
	AuthRoutes routes.Register
	R          *chi.Mux
}

func NewRouter(authRouter routes.Register) *Route {
	return &Route{
		AuthRoutes: authRouter,
		R:          chi.NewRouter(),
	}
}

func (r *Route) Run(mdlwere mv.Middleware) *chi.Mux {
	// Middleware
	r.R.Use(mdlwere.WithLogger, mdlwere.WithAuth)

	// Api
	r.R.Route("/", func(route chi.Router) {
		route.Route("/api/", func(route chi.Router) {
			route.Route("/v1/", func(route chi.Router) {
				route.Route("/auth", r.AuthRoutes.RegisterRoutes) // Аутентификация

				// r.Route("/user") // Пользователь
				// r.Route("/food")   // Обработка блюд
				// r.Route("/ingred") // Обработка ингредиентов
			})
		})
	})
	return r.R
}
