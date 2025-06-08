package routes

import (
	"github.com/genuinebnt/blogify/internal/users/domain/service"
	"github.com/genuinebnt/blogify/internal/users/handler"
	"github.com/genuinebnt/blogify/internal/users/infrastructure/postgres"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func UserRouter(db *pgxpool.Pool) chi.Router {
	r := chi.NewRouter()

	userRepo := postgres.NewPostgresUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r.Post("/auth/register", userHandler.Register())

	return r
}
