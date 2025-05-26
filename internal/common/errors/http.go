package errors

import (
	"fmt"
	"net/http"

	helpers "github.com/genuinebnt/blogify/internal/common"
	"github.com/rs/zerolog/log"
)

func LogError(r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	log.Error().Err(err).Str("method", method).Str("uri", uri).Msg("Request error")
}

func ErrorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := helpers.Envelope{"error": message}
	err := helpers.WriteJSON(w, status, env, nil)
	if err != nil {
		LogError(r, err)
		w.WriteHeader(500)
	}
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	LogError(r, err)
	message := "the server encountered an error and could not process your request"
	ErrorResponse(w, r, http.StatusInternalServerError, message)
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	ErrorResponse(w, r, http.StatusNotFound, message)
}

func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	ErrorResponse(w, r, http.StatusMethodNotAllowed, message)
}
