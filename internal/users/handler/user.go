package handler

import (
	"net/http"

	"github.com/genuinebnt/blogify/internal/common/errors"
	"github.com/genuinebnt/blogify/internal/common/helpers"
	"github.com/genuinebnt/blogify/internal/common/validator"
	"github.com/genuinebnt/blogify/internal/users/domain/entity"
	"github.com/genuinebnt/blogify/internal/users/domain/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return UserHandler{
		userService: userService,
	}
}

func (u *UserHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := entity.User{}
		err := helpers.ReadJSON(w, r, &user)
		if err != nil {
			errors.BadRequestResponse(w, r, err)
			return
		}

		v := validator.New()
		if user.Validate(v); !v.Valid() {
			errors.FailedValidationResponse(w, r, v.Errors)
			return
		}

		u.userService.CreateUser(&user)
	}
}
