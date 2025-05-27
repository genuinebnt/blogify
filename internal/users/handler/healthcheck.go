package handler

import (
	"net/http"

	"github.com/genuinebnt/blogify/internal/common/config"
	"github.com/genuinebnt/blogify/internal/common/errors"
	"github.com/genuinebnt/blogify/internal/common/helpers"
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
	env := helpers.Envelope{
		"status": "available",
		"system_info": map[string]string{
			"environmen": h.cfg.Env,
			"version":    version,
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		err := helpers.WriteJSON(w, http.StatusOK, env, nil)
		errors.ServerErrorResponse(w, r, err)
	}
}
