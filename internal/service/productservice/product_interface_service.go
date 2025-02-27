package productService

import (
	"context"

	"github.com/DevJonathanSantos/poc-go-api/internal/dto"
	"github.com/DevJonathanSantos/poc-go-api/internal/handler/response"
	productRepository "github.com/DevJonathanSantos/poc-go-api/internal/repository/productrepository"
)

func NewProductService(repo productRepository.ProductRepository) ProductService {
	return &service{
		repo,
	}
}

type service struct {
	repo productRepository.ProductRepository
}

type ProductService interface {
	CreateProduct(ctx context.Context, u dto.CreateProductDto) error
	UpdateProduct(ctx context.Context, id string, u dto.UpdateProductDto) error
	DeleteProduct(ctx context.Context, id string) error
	FindManyProducts(ctx context.Context, d dto.FindProductDto) ([]response.ProductResponse, error)
}
