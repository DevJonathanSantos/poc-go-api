package userService

import userRepository "github.com/DevJonathanSantos/poc-go-api/internal/repository/userepository"

func NewUserService(repo userRepository.UserRepository) UserService {
	return &service{
		repo,
	}
}

type service struct {
	repo userRepository.UserRepository
}

type UserService interface {
	CreateUser() error
}
