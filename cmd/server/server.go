package server

import (
	"mealmate/cmd/config"
	"mealmate/internal/logg"
	mv "mealmate/internal/middleware"
	"net/http"
)

type Server struct {
	Cfg  config.Config
	Logg logg.Logger
}

func NewServer(config config.Config, logger logg.Logger) *Server {
	return &Server{Cfg: config, Logg: logger}
}

func (s *Server) Run(router Router, mdlwere mv.Middleware) {
	s.Logg.RaiseFatal(
		"Server is bad",
		http.ListenAndServe(s.Cfg.GetRunAddress(), router.Run(mdlwere)))
}
