package server

import "github.com/go-chi/chi"

type Router interface {
	Run() *chi.Mux
}
