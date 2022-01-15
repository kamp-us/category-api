package backend

import (
	"context"
	"github.com/kamp-us/category-api/internal/models"
)

func (b *PostgreSQLBackend) DeleteCategory(ctx context.Context, id string) error {
	category := models.Category{}
	result := b.DB.First(&category, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	result = b.DB.Delete(&category)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
