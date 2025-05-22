package main

import (
	"github.com/genuinebnt/blogify/internal/users/routes"
	"github.com/go-chi/chi/v5"
)

func Router() chi.Router {
	r := chi.NewRouter()
	r.Mount("/", routes.HealthCheckRouter())
	return r
}
