package userService

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/DevJonathanSantos/poc-go-api/api/viacep"
	"github.com/DevJonathanSantos/poc-go-api/internal/dto"
	"github.com/DevJonathanSantos/poc-go-api/internal/entity"
	"github.com/DevJonathanSantos/poc-go-api/internal/handler/response"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) CreateUser(ctx context.Context, u dto.CreateUserDto) error {

	userExists, err := s.repo.FindUserByEmail(ctx, u.Email)
	if err != nil {
		slog.Error("error to search user by email", "err", err, slog.String("package", "userService"))
		return err
	}
	if userExists != nil {
		slog.Error("user already exists", slog.String("package", "userService"))
		return errors.New("user already exists")
	}

	passwordEncrypted, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		slog.Error("error to encrypt password", "err", err, slog.String("package", "userService"))
		return errors.New("error to encrypt password")
	}

	cep, err := viacep.GetCep(u.CEP)
	if err != nil {
		slog.Error("error to get cep", "err", err, slog.String("package", "userservice"))
		return err
	}
	newUser := entity.UserEntity{
		ID:       uuid.New().String(),
		Name:     u.Name,
		Email:    u.Email,
		Password: string(passwordEncrypted),
		Address: entity.UserAddress{
			CEP:        cep.CEP,
			IBGE:       cep.IBGE,
			UF:         cep.UF,
			City:       cep.Localidade,
			Complement: cep.Complemento,
			Street:     cep.Logradouro,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.repo.CreateUser(ctx, &newUser)
	if err != nil {
		slog.Error("error to create user", "err", err, slog.String("package", "userService"))
		return err
	}

	return nil
}

func (s *service) UpdateUser(ctx context.Context, u dto.UpdateUserDto, id string) error {

	userExists, err := s.repo.FindUserByID(ctx, id)
	if err != nil {
		slog.Error("error to search user by id", "err", err, slog.String("package", "userService"))
		return err
	}
	if userExists == nil {
		slog.Error("user not found", slog.String("package", "userService"))
		return errors.New("user already exists")
	}

	var updateUser entity.UserEntity
	if u.Email != "" {
		verifyUserEmail, err := s.repo.FindUserByEmail(ctx, u.Email)
		if err != nil {
			slog.Error("error to search user by email", "err", err, slog.String("package", "userservice"))
			return err
		}
		if verifyUserEmail != nil {
			slog.Error("user already exists", slog.String("package", "userservice"))
			return errors.New("user already exists")
		}
		updateUser.Email = u.Email
	}
	if u.CEP != "" {
		cep, err := viacep.GetCep(u.CEP)
		if err != nil {
			slog.Error("error to get cep", "err", err, slog.String("package", "userservice"))
			return err
		}
		updateUser.Address = entity.UserAddress{
			CEP:        cep.CEP,
			IBGE:       cep.IBGE,
			UF:         cep.UF,
			City:       cep.Localidade,
			Complement: cep.Complemento,
			Street:     cep.Logradouro,
		}
	}

	updateUser.ID = id
	updateUser.Name = u.Name
	updateUser.UpdatedAt = time.Now()
	err = s.repo.UpdateUser(ctx, &updateUser)
	if err != nil {
		slog.Error("error to update user", "err", err, slog.String("package", "userService"))
		return err
	}
	return nil
}

func (s *service) GetUserByID(ctx context.Context, id string) (*response.UserResponse, error) {
	userExists, err := s.repo.FindUserByID(ctx, id)
	if err != nil {
		slog.Error("error to search user by id", "err", err, slog.String("package", "userService"))
		return nil, err
	}

	if userExists == nil {
		slog.Error("user not found", slog.String("package", "userService"))
		return nil, errors.New("user not found")
	}

	user := response.UserResponse{
		ID:        userExists.ID,
		Name:      userExists.Name,
		Email:     userExists.Email,
		CreatedAt: userExists.CreatedAt,
		UpdatedAt: userExists.UpdatedAt,
	}
	return &user, nil
}

func (s *service) DeleteUser(ctx context.Context, id string) error {

	slog.Info("OOOOOOOOOOOOOOOO")
	userExists, err := s.repo.FindUserByID(ctx, id)
	if err != nil {
		slog.Error("error to search user by id", "err", err, slog.String("package", "userService"))
		return err
	}
	if userExists == nil {
		slog.Error("user not found", slog.String("package", "userService"))
		return errors.New("user not found")
	}
	err = s.repo.DeleteUser(ctx, id)
	if err != nil {
		slog.Error("error to delete user", "err", err, slog.String("package", "userService"))
		return err
	}
	return nil
}

func (s *service) FindManyUsers(ctx context.Context) (*response.ManyUsersResponse, error) {
	findManyUsers, err := s.repo.FindManyUsers(ctx)
	if err != nil {
		slog.Error("error to find many users", "err", err, slog.String("package", "userService"))
		return nil, err
	}
	users := response.ManyUsersResponse{}
	for _, user := range findManyUsers {
		userResponse := response.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		users.Users = append(users.Users, userResponse)
	}
	return &users, nil
}

func (s *service) UpdateUserPassword(ctx context.Context, u *dto.UpdateUserPasswordDto, id string) error {
	userExists, err := s.repo.FindUserByID(ctx, id)
	if err != nil {
		slog.Error("error to search user by id", "err", err, slog.String("package", "userservice"))
		return err
	}
	if userExists == nil {
		slog.Error("user not found", slog.String("package", "userservice"))
		return errors.New("user not found")
	}
	oldPass, err := s.repo.GetUserPassword(ctx, id)
	if err != nil {
		slog.Error("error to get user password", "err", err, slog.String("package", "userservice"))
		return err
	}
	// compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(oldPass), []byte(u.OldPassword))
	if err != nil {
		slog.Error("invalid password", slog.String("package", "userservice"))
		return errors.New("invalid password")
	}
	// compare new password with password in database
	err = bcrypt.CompareHashAndPassword([]byte(oldPass), []byte(u.Password))
	if err == nil {
		slog.Error("new password is equal to old password", slog.String("package", "userservice"))
		return errors.New("new password is equal to old password")
	}
	passwordEncrypted, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		slog.Error("error to encrypt password", "err", err, slog.String("package", "userservice"))
		return errors.New("error to encrypt password")
	}
	err = s.repo.UpdatePassword(ctx, string(passwordEncrypted), id)
	if err != nil {
		slog.Error("error to update password", "err", err, slog.String("package", "userservice"))
		return err
	}
	return nil
}
