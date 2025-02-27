package userService

import (
	"context"

	"github.com/DevJonathanSantos/poc-go-api/internal/dto"
)

func (s *service) CreateUser(ctx context.Context, u dto.CreateUserDto) error {
	return nil
}

func (s *service) UpdateUser(ctx context.Context, u dto.UpdateUserDto, id string) error {
	return nil
}
