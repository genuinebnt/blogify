package repository

import "github.com/genuinebnt/blogify/internal/users/domain/entity"

type UserRepository interface {
	Insert(entity.User) error
	Update(entity.User) error
	Delete(id int64) error
	FindAll() error
	FindByID(d int64) error
	FindByEmail(email string) error
}
