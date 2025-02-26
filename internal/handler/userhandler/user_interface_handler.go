package user_handler

import "net/http"

func NewUserHandler() UserHandler {
	return &handler{}
}

type handler struct {
}

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}
