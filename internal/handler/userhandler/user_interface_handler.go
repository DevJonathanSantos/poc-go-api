package userHandler

import (
	"net/http"

	userService "github.com/DevJonathanSantos/poc-go-api/internal/service/userservice"
)

func NewUserHandler(service userService.UserService) UserHandler {
	return &handler{
		service,
	}
}

type handler struct {
	service userService.UserService
}

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	FindManyUsers(w http.ResponseWriter, r *http.Request)
	UpdateUserPassword(w http.ResponseWriter, r *http.Request)
}
