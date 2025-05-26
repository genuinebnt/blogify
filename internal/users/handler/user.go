package handler

import (
	"net/http"

	"github.com/genuinebnt/blogify/internal/users/domain/service"
	"github.com/rs/zerolog/log"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return UserHandler{
		userService: userService,
	}
}

func (u *UserHandler) ListUsers() http.HandlerFunc {
	err := u.userService.UserRepo.FindAll()
	if err != nil {
		log.Error().Msg(err.Error())
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
