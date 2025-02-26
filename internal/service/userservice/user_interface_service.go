package userservice

func NewUserService() UserService {
	return &service{}
}

type service struct {
}

type UserService interface {
	CreateUser() error
}
