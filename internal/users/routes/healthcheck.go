package routes

import (
	"github.com/genuinebnt/blogify/internal/users/handler"
	"github.com/go-chi/chi/v5"
)

func HealthCheckRouter() chi.Router {
	r := chi.NewRouter()

	healthcheckHandler := handler.NewHealthCheckHandler()
	r.Get("/healthcheck", healthcheckHandler.CheckHealth())

	return r
}
