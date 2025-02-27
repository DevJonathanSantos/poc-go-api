package categoryService

import (
	"context"
	"errors"
	"time"

	"github.com/DevJonathanSantos/poc-go-api/internal/dto"
	"github.com/DevJonathanSantos/poc-go-api/internal/entity"
	"github.com/google/uuid"
)

func (s *service) CreateCategory(ctx context.Context, u dto.CreateCategoryDto) error {
	categoryEntity := entity.CategoryEntity{
		ID:        uuid.New().String(),
		Title:     u.Title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := s.repo.CreateCategory(ctx, &categoryEntity)
	if err != nil {
		return errors.New("error to create category")
	}
	return nil
}
