package server

import (
	"context"
	"errors"
	"net/http"
	"sync"

	"github.com/boginskiy/FoodMoment/foodservice/internal/logg"
)

type HTTPServer struct {
	S       *http.Server
	Logger  logg.Logger
	runOnce sync.Once
}

func NewHTTPServer(ctx context.Context, cfg Config, log logg.Logger) (*HTTPServer, error) {
	server := &http.Server{
		Addr:         cfg.Addr,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
	return &HTTPServer{
		S:      server,
		Logger: log,
	}, nil
}

func (s *HTTPServer) Run(ctx context.Context, handler http.Handler) error {
	if handler == nil {
		return errors.New("handler cannot be nil")
	}
	s.S.Handler = handler

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), defaultShutdownTimeout)
		defer cancel()
		if err := s.S.Shutdown(shutdownCtx); err != nil {
			s.Logger.Error("server shutdown error",
				"struct", "HTTPServer",
				"error", err)
		}
	}()

	// Start
	err := s.S.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
