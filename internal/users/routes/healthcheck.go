package routes

import (
	"github.com/genuinebnt/blogify/internal/users/handlers"
	"github.com/go-chi/chi/v5"
)

func HealthCheckRouter() chi.Router {
	r := chi.NewRouter()

	healthcheckHandler := handlers.NewHealthCheckHandler()
	r.Get("/healthcheck", healthcheckHandler.CheckHealth())

	return r
}
