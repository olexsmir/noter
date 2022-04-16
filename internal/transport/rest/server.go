package rest

import (
	"context"
	"net/http"

	"github.com/olexsmir/noter/internal/config"
)

type Server struct {
	http *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		http: &http.Server{
			Addr:           ":" + cfg.HTTP.Port,
			Handler:        handler,
			ReadTimeout:    cfg.HTTP.ReadTimeout,
			WriteTimeout:   cfg.HTTP.WriteTimeout,
			MaxHeaderBytes: cfg.HTTP.MaxHeaderMegabytes << 20,
		},
	}
}

func (s *Server) Start() error {
	return s.http.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}
