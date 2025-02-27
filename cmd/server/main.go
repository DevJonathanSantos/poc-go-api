package main

import (
	"fmt"
	"log/slog"
	"net/http"

	logger "github.com/DevJonathanSantos/poc-go-api/config"
	"github.com/DevJonathanSantos/poc-go-api/config/env"
	_ "github.com/DevJonathanSantos/poc-go-api/docs"
	"github.com/DevJonathanSantos/poc-go-api/internal/database"
	"github.com/DevJonathanSantos/poc-go-api/internal/database/sqlc"
	"github.com/DevJonathanSantos/poc-go-api/internal/handler"
	"github.com/DevJonathanSantos/poc-go-api/internal/handler/routes"

	"github.com/go-chi/chi"

	categoryRepository "github.com/DevJonathanSantos/poc-go-api/internal/repository/categoryrepository"
	productRepository "github.com/DevJonathanSantos/poc-go-api/internal/repository/productrepository"
	userRepository "github.com/DevJonathanSantos/poc-go-api/internal/repository/userepository"

	categoryService "github.com/DevJonathanSantos/poc-go-api/internal/service/categoryservice"
	productService "github.com/DevJonathanSantos/poc-go-api/internal/service/productservice"
	userService "github.com/DevJonathanSantos/poc-go-api/internal/service/userservice"
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

	queries := sqlc.New(dbConnection)

	// user
	userRepo := userRepository.NewUserRepository(dbConnection, queries)
	newUserService := userService.NewUserService(userRepo)

	// category
	categoryRepo := categoryRepository.NewCategoryRepository(dbConnection, queries)
	newCategoryService := categoryService.NewCategoryService(categoryRepo)

	// product
	productRepo := productRepository.NewProductRepository(dbConnection, queries)
	productsService := productService.NewProductService(productRepo)

	newHandler := handler.NewHandler(newUserService, newCategoryService, productsService)

	// init routes
	router := chi.NewRouter()
	routes.InitRoutes(router, newHandler)
	routes.InitDocsRoutes(router)

	port := fmt.Sprintf(":%s", env.Env.GoPort)
	slog.Info(fmt.Sprintf("server running on port %s", port))
	err = http.ListenAndServe(port, router)
	if err != nil {
		slog.Error("error to start server", err, slog.String("package", "main"))
	}
}
