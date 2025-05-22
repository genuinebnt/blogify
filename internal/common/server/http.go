package server

import (
	"net/http"
	"os"

	"github.com/genuinebnt/blogify/internal/common/logs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

func RunHTTPServer(router chi.Router) {
	RunHTTPServerOnAddr(":"+os.Getenv("PORT"), router)
}

func RunHTTPServerOnAddr(addr string, router chi.Router) {
	rootRouter := chi.NewRouter()
	setMiddlewares(rootRouter)
	rootRouter.Mount("/api/v1/", router)

	err := http.ListenAndServe(addr, rootRouter)
	if err != nil {
		logs.GetLogger().Error().Err(err).Str("addr", addr).Msg("HTTP server failed to start")
	}

	logs.GetLogger().Info().Msg("Server started")
}

func setMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	router.Use(logs.NewStructuredLogger(&logs.ZeroLogLogger{Logger: logger}))
}
