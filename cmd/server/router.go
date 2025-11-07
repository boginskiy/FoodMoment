package server

import "github.com/go-chi/chi"

type Route struct {
	R *chi.Mux
}

func NewRouter() *Route {
	return &Route{
		R: chi.NewRouter(),
	}
}

func (r *Route) Run() {
	// Middleware

	// Api
	r.R.Route("/", func(r chi.Router) {
		r.Route("/api/", func(r chi.Router) {
			r.Route("/v1/", func(r chi.Router) {
				r.Route("/auth")   // Аутентификация
				r.Route("/user")   // Пользователь
				r.Route("/food")   // Обработка блюд
				r.Route("/ingred") // Обработка ингредиентов
			})
		})
	})
}
