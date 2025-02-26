package main

import (
	"fmt"
	"log/slog"
	"net/http"

	logger "github.com/DevJonathanSantos/poc-go-api/config"
	"github.com/DevJonathanSantos/poc-go-api/config/env"
	"github.com/DevJonathanSantos/poc-go-api/internal/database"
	"github.com/DevJonathanSantos/poc-go-api/internal/database/sqlc"
	"github.com/DevJonathanSantos/poc-go-api/internal/handler/routes"
	userHandler "github.com/DevJonathanSantos/poc-go-api/internal/handler/userhandler"
	userRepository "github.com/DevJonathanSantos/poc-go-api/internal/repository/userepository"
	userService "github.com/DevJonathanSantos/poc-go-api/internal/service/userservice"
	"github.com/go-chi/chi"
)

func main() {
	logger.InitLogger()
	slog.Info("starting api")

	_, err := env.LoadingConfig(".")
	if err != nil {
		slog.Error("failed to load environment variables", err, slog.String("package", "main"))
		return
	}
	dbConnection, err := database.NewDBConnection()
	if err != nil {
		slog.Error("error to connect to database", "err", err, slog.String("package", "main"))
		return
	}

	router := chi.NewRouter()
	queries := sqlc.New(dbConnection)

	// user
	userRepo := userRepository.NewUserRepository(dbConnection, queries)
	newUserService := userService.NewUserService(userRepo)
	newUserHandler := userHandler.NewUserHandler(newUserService)

	// init routes
	routes.InitUserRoutes(router, newUserHandler)

	port := fmt.Sprintf(":%s", env.Env.GoPort)
	slog.Info(fmt.Sprintf("server running on port %s", port))
	err = http.ListenAndServe(port, router)
	if err != nil {
		slog.Error("error to start server", err, slog.String("package", "main"))
	}
}
