package logs

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
)

func NewStructuredLogger(formatter middleware.LogFormatter) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(formatter)
}
