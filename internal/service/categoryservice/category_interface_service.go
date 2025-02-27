package categoryService

import (
	"context"

	"github.com/DevJonathanSantos/poc-go-api/internal/dto"
	categoryRepository "github.com/DevJonathanSantos/poc-go-api/internal/repository/categoryrepository"
)

func NewCategoryService(repo categoryRepository.CategoryRepository) CategoryService {
	return &service{
		repo,
	}
}

type service struct {
	repo categoryRepository.CategoryRepository
}

type CategoryService interface {
	CreateCategory(ctx context.Context, u dto.CreateCategoryDto) error
}
