package repository

import (
	"github.com/genuinebnt/blogify/internal/users/domain/entity"
	"github.com/google/uuid"
)

type UserRepository interface {
	Insert(*entity.User) error
	Update(*entity.User) error
	Delete(id int64) error
	FindAll() error
	FindByID(id uuid.UUID) error
	FindByEmail(email string) error
}
