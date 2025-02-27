package routes

import (
	"github.com/DevJonathanSantos/poc-go-api/config/env"
	"github.com/DevJonathanSantos/poc-go-api/internal/handler/middleware"
	userHandler "github.com/DevJonathanSantos/poc-go-api/internal/handler/userhandler"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

func InitUserRoutes(router chi.Router, h userHandler.UserHandler) {
	router.Use(middleware.LoggerData)

	router.Post("/user", h.CreateUser)
	router.Route("/user", func(r chi.Router) {
		r.Use(jwtauth.Verifier(env.Env.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Patch("/", h.UpdateUser)
		r.Get("/", h.GetUserByID)
		r.Delete("/", h.DeleteUser)
		r.Get("/list-all", h.FindManyUsers)
		r.Patch("/password", h.UpdateUserPassword)
	})
	router.Route("/auth", func(r chi.Router) {
		r.Post("/login", h.Login)
	})
}
