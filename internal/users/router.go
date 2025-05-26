package main

import (
	"github.com/genuinebnt/blogify/internal/common/config"
	"github.com/genuinebnt/blogify/internal/users/routes"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GlobalRouter(db *pgxpool.Pool, cfg *config.Config) chi.Router {
	r := chi.NewRouter()
	r.Mount("/", routes.HealthCheckRouter(cfg))
	r.Mount("/users", routes.UserRouter(db))
	return r
}
