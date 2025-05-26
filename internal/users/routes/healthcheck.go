package routes

import (
	"github.com/genuinebnt/blogify/internal/common/config"
	"github.com/genuinebnt/blogify/internal/users/handler"
	"github.com/go-chi/chi/v5"
)

func HealthCheckRouter(cfg *config.Config) chi.Router {
	r := chi.NewRouter()

	healthcheckHandler := handler.NewHealthCheckHandler(cfg)
	r.Get("/healthcheck", healthcheckHandler.CheckHealth())

	return r
}
