package handler

import (
	"net/http"

	categoryService "github.com/DevJonathanSantos/poc-go-api/internal/service/categoryservice"
	productService "github.com/DevJonathanSantos/poc-go-api/internal/service/productservice"
	userService "github.com/DevJonathanSantos/poc-go-api/internal/service/userservice"
)

func NewHandler(userService userService.UserService,
	categoryService categoryService.CategoryService,
	productservice productService.ProductService) Handler {
	return &handler{
		userService:     userService,
		categoryService: categoryService,
		productservice:  productservice,
	}
}

type handler struct {
	userService     userService.UserService
	categoryService categoryService.CategoryService
	productservice  productService.ProductService
}

type Handler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	FindManyUsers(w http.ResponseWriter, r *http.Request)
	UpdateUserPassword(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)

	CreateCategory(w http.ResponseWriter, r *http.Request)

	CreateProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
	FindManyProducts(w http.ResponseWriter, r *http.Request)
}
