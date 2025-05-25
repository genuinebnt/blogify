package main

import (
	"github.com/genuinebnt/blogify/internal/common/config"
	postgres "github.com/genuinebnt/blogify/internal/common/db"
	"github.com/genuinebnt/blogify/internal/common/logs"
	"github.com/genuinebnt/blogify/internal/common/server"
	"github.com/rs/zerolog/log"
)

func main() {
	// Initialize the global logger
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Error().Msgf("Failed to load config %s", err.Error())
		return
	}

	logs.Init(cfg)
	db, err := postgres.NewPostgresDB(cfg.ConnectionStringFromEnv())
	if err != nil {
		log.Error().Msg("Failed to connect to postgres")
		return
	}

	r := Router()
	server.RunHTTPServer(r, cfg)
}
