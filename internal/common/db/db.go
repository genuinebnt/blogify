package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

func NewPostgresDB(connection string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbPool, err := pgxpool.New(ctx, connection)
	if err != nil {
		return nil, err
	}

	log.Info().Msg("Database connection pool established")

	return dbPool, nil
}
