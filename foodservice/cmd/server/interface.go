package server

import (
	mv "mealmate/internal/middleware"

	"github.com/go-chi/chi"
)

type Router interface {
	Run(mdlwere mv.Middleware) *chi.Mux
}
