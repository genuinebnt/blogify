package service

import "github.com/genuinebnt/blogify/internal/users/domain/repository"

type UserService struct {
	UserRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return UserService{
		UserRepo: userRepo,
	}
}
