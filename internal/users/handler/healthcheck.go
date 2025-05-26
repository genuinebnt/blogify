package handler

import (
	"net/http"

	helpers "github.com/genuinebnt/blogify/internal/common"
	"github.com/genuinebnt/blogify/internal/common/config"
)

const version = "0.1.0"

type HealthCheckHandler struct {
	cfg *config.Config
}

func NewHealthCheckHandler(cfg *config.Config) *HealthCheckHandler {
	return &HealthCheckHandler{
		cfg,
	}
}

func (h *HealthCheckHandler) CheckHealth() http.HandlerFunc {
	data := map[string]string{
		"status":      "available",
		"environment": h.cfg.Env,
		"version":     version,
	}

	return func(w http.ResponseWriter, r *http.Request) {
		err := helpers.WriteJSON(w, http.StatusOK, data, nil)
		if err != nil {
			http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		}
	}
}
