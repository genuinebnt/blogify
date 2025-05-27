package service

import (
	"time"

	"github.com/genuinebnt/blogify/internal/users/domain/entity"
	"github.com/genuinebnt/blogify/internal/users/domain/repository"
	"github.com/google/uuid"
)

type UserService struct {
	UserRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return UserService{
		UserRepo: userRepo,
	}
}

func (u UserService) CreateUser(user *entity.User) {
	user.Id = uuid.New()
	user.CreatedAt = time.Now()

	u.UserRepo.Create(user)
}
