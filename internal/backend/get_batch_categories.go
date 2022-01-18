package backend

import (
	"context"
	"github.com/kamp-us/category-api/internal/models"
)

func (b *PostgreSQLBackend) GetBatchCategories(ctx context.Context, ids []string) ([]*models.Category, error) {
	var categories []*models.Category
	result := b.DB.Find(&categories, ids)
	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}
