package repository

import (
	"github.com/genuinebnt/blogify/internal/users/domain/entity"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(*entity.User) error
	Update(*entity.User) error
	Delete(id int64) error
	FindAll() ([]entity.User, error)
	FindByID(id uuid.UUID) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}
