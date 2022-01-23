package backend

import (
	"context"
	"github.com/kamp-us/category-api/internal/models"
)

func (b *PostgreSQLBackend) GetCategoryBySlug(ctx context.Context, slug string) (*models.Category, error) {
	category := models.Category{}
	result := b.DB.First(&category, "slug = ?", slug)
	if result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}
