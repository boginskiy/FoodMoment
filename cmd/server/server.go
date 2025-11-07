package server

import (
	"mealmate/cmd/config"
	"mealmate/internal/logg"
	"net/http"
)

type Server struct {
	Cfg  config.Config
	Logg logg.Logger
}

func NewServer(config config.Config, logger logg.Logger) *Server {
	return &Server{Cfg: config, Logg: logger}
}

func (s *Server) Run() {
	s.Logg.RaiseFatal(
		"Server is bad",
		http.ListenAndServe(s.Cfg.GetRunAddress(), r.Router()))
}
