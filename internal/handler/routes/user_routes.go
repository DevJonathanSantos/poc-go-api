package routes

import (
	userHandler "github.com/DevJonathanSantos/poc-go-api/internal/handler/userhandler"
	"github.com/go-chi/chi"
)

func InitUserRoutes(router chi.Router, h userHandler.UserHandler) {
	router.Route("/user", func(r chi.Router) {
		r.Post("/", h.CreateUser)
		r.Patch("/{id}", h.UpdateUser)
		r.Get("/{id}", h.GetUserByID)
		r.Delete("/{id}", h.DeleteUser)
		r.Get("/", h.FindManyUsers)
	})
}
