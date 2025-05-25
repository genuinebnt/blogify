package postgres

import (
	"github.com/genuinebnt/blogify/internal/users/domain/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUserRepository struct {
	db *pgxpool.Pool
}

func NewPostgresUserRepository(db *pgxpool.Pool) *PostgresUserRepository {
	return &PostgresUserRepository{db}
}

func (u *PostgresUserRepository) Insert(entity.User) error {
	query := `
		INSERT INTO users (id, username, email, password)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at`

}
