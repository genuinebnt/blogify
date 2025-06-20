package main

import (
	"github.com/genuinebnt/blogify/internal/common/config"
	postgres "github.com/genuinebnt/blogify/internal/common/db"
	"github.com/genuinebnt/blogify/internal/common/logs"
	"github.com/genuinebnt/blogify/internal/common/server"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.LoadConfig()
	log.Info().Msgf("Loaded config: %+v", cfg)
	if err != nil {
		log.Error().Msgf("Failed to load config %s", err.Error())
		return
	}

	logs.Init(cfg)
	db, err := postgres.NewPostgresDB(cfg.ConnectionStringFromEnv())
	if err != nil {
		log.Error().Msgf("Failed to connect to postgres with error: %s", err)
		return
	}

	r := GlobalRouter(db, cfg)
	s := server.NewServer(r, cfg)
	s.RunHTTPServer()
}
