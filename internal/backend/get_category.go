package backend

import (
	"context"
	"github.com/kamp-us/category-api/internal/models"
)

func (b *PostgreSQLBackend) GetCategory(ctx context.Context, id string) (*models.Category, error) {
	category := models.Category{}
	result := b.DB.First(&category, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}
