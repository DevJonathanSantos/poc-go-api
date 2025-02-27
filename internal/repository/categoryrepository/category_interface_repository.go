package categoryrepository

import (
	"context"
	"database/sql"

	"github.com/DevJonathanSantos/poc-go-api/internal/database/sqlc"
	"github.com/DevJonathanSantos/poc-go-api/internal/entity"
)

func NewCategoryRepository(db *sql.DB, q *sqlc.Queries) CategoryRepository {
	return &repository{
		db,
		q,
	}
}

type repository struct {
	db      *sql.DB
	queries *sqlc.Queries
}

type CategoryRepository interface {
	CreateCategory(ctx context.Context, c *entity.CategoryEntity) error
}
