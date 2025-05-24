package server

import (
	"net/http"
	"os"

	"github.com/genuinebnt/blogify/internal/common/logs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

func RunHTTPServer(router chi.Router) {
	RunHTTPServerOnAddr(os.Getenv("PORT"), router)
}

func RunHTTPServerOnAddr(port string, router chi.Router) {
	rootRouter := chi.NewRouter()
	setMiddlewares(rootRouter)
	rootRouter.Mount("/api/v1/", router)

	log.Info().Msgf("Starting server on port: %s", port)
	err := http.ListenAndServe(":"+port, rootRouter)

	if err != nil {
		log.Error().Msgf("Failed to start server with err: %s", err.Error())
	}
}

func setMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)

	router.Use(logs.NewStructuredLogger(&logs.ZeroLogLogger{Logger: log.Logger}))
}
