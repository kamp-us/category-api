package backend

import (
	"context"
	"github.com/gosimple/slug"
	"github.com/kamp-us/category-api/internal/models"
)

func (b *PostgreSQLBackend) CreateCategory(ctx context.Context, name string, description string) (*models.Category, error) {
	category := models.Category{
		Slug:        slug.MakeLang(name, "tr"),
		Name:        name,
		Description: description,
	}

	result := b.DB.Create(&category)
	if result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}
