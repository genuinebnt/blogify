package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/genuinebnt/blogify/internal/common/config"
	"github.com/genuinebnt/blogify/internal/common/errors"
	"github.com/genuinebnt/blogify/internal/common/logs"
	custom_middleware "github.com/genuinebnt/blogify/internal/common/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

type Server struct {
	httpServer *http.Server
	cfg        *config.Config
}

func NewServer(router chi.Router, cfg *config.Config) *Server {
	rootRouter := chi.NewRouter()

	rootRouter.NotFound(errors.NotFoundResponse)
	rootRouter.MethodNotAllowed(errors.MethodNotAllowedResponse)

	setMiddlewares(rootRouter)
	rootRouter.Mount("/api/v1/", router)

	return &Server{
		httpServer: &http.Server{
			Addr:         fmt.Sprintf(":%d", cfg.Port),
			Handler:      rootRouter,
			IdleTimeout:  time.Minute,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}, cfg: cfg,
	}
}

func (s *Server) RunHTTPServer() {
	log.Info().Msgf("Starting server on port: %d", s.cfg.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", s.cfg.Port), s.httpServer.Handler)

	if err != nil {
		log.Error().Msgf("Failed to start server with err: %s", err.Error())
	}
}

func setMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(logs.NewStructuredLogger(&logs.ZeroLogLogger{Logger: log.Logger}))
	router.Use(custom_middleware.RecoverPanic)
}
