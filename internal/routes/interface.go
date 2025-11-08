package routes

import "github.com/go-chi/chi"

type Register interface {
	RegisterRoutes(route chi.Router)
}
