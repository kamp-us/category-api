package backend

import (
	"context"
	"github.com/kamp-us/category-api/internal/models"
	"gorm.io/gorm"
)

type Backender interface {
	GetCategory(ctx context.Context, id string) (*models.Category, error)
	CreateCategory(ctx context.Context, name string, description string) (*models.Category, error)
	UpdateCategory(ctx context.Context, id string, name *string, description *string) error
	DeleteCategory(ctx context.Context, id string) error
	GetBatchCategories(ctx context.Context, ids []string) ([]*models.Category, error)
	GetCategoryBySlug(ctx context.Context, slug string) (*models.Category, error)
}

type PostgreSQLBackend struct {
	DB *gorm.DB
}

func NewPostgreSQLBackend(db *gorm.DB) Backender {
	return &PostgreSQLBackend{
		DB: db,
	}
}
