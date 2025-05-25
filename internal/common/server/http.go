package server

import (
	"fmt"
	"net/http"

	"github.com/genuinebnt/blogify/internal/common/config"
	"github.com/genuinebnt/blogify/internal/common/logs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

func RunHTTPServer(router chi.Router, cfg *config.Config) {
	RunHTTPServerOnAddr(cfg.Port, router)
}

func RunHTTPServerOnAddr(port int64, router chi.Router) {
	rootRouter := chi.NewRouter()
	setMiddlewares(rootRouter)
	rootRouter.Mount("/api/v1/", router)

	log.Info().Msgf("Starting server on port: %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), rootRouter)

	if err != nil {
		log.Error().Msgf("Failed to start server with err: %s", err.Error())
	}
}

func setMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)

	router.Use(logs.NewStructuredLogger(&logs.ZeroLogLogger{Logger: log.Logger}))
}
