package userService

import (
	"context"

	"github.com/DevJonathanSantos/poc-go-api/internal/dto"
	"github.com/DevJonathanSantos/poc-go-api/internal/handler/response"
	userRepository "github.com/DevJonathanSantos/poc-go-api/internal/repository/userepository"
)

func NewUserService(repo userRepository.UserRepository) UserService {
	return &service{
		repo,
	}
}

type service struct {
	repo userRepository.UserRepository
}

type UserService interface {
	CreateUser(ctx context.Context, u dto.CreateUserDto) error
	UpdateUser(ctx context.Context, u dto.UpdateUserDto, id string) error
	GetUserByID(ctx context.Context, id string) (*response.UserResponse, error)
	DeleteUser(ctx context.Context, id string) error
	FindManyUsers(ctx context.Context) (*response.ManyUsersResponse, error)
	UpdateUserPassword(ctx context.Context, u *dto.UpdateUserPasswordDto, id string) error
}
