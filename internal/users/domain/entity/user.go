package entity

import (
	"time"

	"github.com/genuinebnt/blogify/internal/common/validator"
	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"-"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"-"`
}

func (u User) Validate(v *validator.Validator) {
	v.Check(u.Username != "", "username", "must be provided")
	v.Check(len(u.Username) <= 64, "username", "must not be more than 64 characters")

	v.Check(u.Password != "", "password", "must be provided")
	v.Check(len(u.Password) <= 8, "password", "must not be more than 8 characters")
}
