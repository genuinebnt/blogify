package postgres

import (
	"context"

	"github.com/genuinebnt/blogify/internal/users/domain/entity"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUserRepository struct {
	db *pgxpool.Pool
}

func NewPostgresUserRepository(db *pgxpool.Pool) PostgresUserRepository {
	return PostgresUserRepository{
		db: db,
	}
}

func (u PostgresUserRepository) Insert(user *entity.User) error {
	query := `
		INSERT INTO users (id, username, email, password)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at`

	args := []any{user.Id, user.Username, user.Email, user.Password}
	return u.db.QueryRow(context.Background(), query, args...).Scan(&user.Id, &user.CreatedAt)
}

func (u PostgresUserRepository) Update(user *entity.User) error {
	return nil
}

func (u PostgresUserRepository) Delete(id int64) error {
	return nil
}

func (u PostgresUserRepository) FindAll() error {
	return nil
}

func (u PostgresUserRepository) FindByID(id uuid.UUID) error {
	return nil
}

func (u PostgresUserRepository) FindByEmail(email string) error {
	return nil
}
