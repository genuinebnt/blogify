package main

import (
	"github.com/genuinebnt/blogify/internal/users/routes"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GlobalRouter(db *pgxpool.Pool) chi.Router {
	r := chi.NewRouter()
	r.Mount("/", routes.HealthCheckRouter())
	r.Mount("/users", routes.UserRouter(db))
	return r
}
