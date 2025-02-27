package userService

import (
	"context"
	"time"

	"github.com/DevJonathanSantos/poc-go-api/internal/dto"
	"github.com/DevJonathanSantos/poc-go-api/internal/handler/response"
)

func (s *service) CreateUser(ctx context.Context, u dto.CreateUserDto) error {
	return nil
}

func (s *service) UpdateUser(ctx context.Context, u dto.UpdateUserDto, id string) error {
	return nil
}

func (s *service) GetUserByID(ctx context.Context, id string) (*response.UserResponse, error) {
	userFake := response.UserResponse{
		ID:        "123",
		Name:      "John Doe",
		Email:     "jonh.doe@email.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return &userFake, nil
}
