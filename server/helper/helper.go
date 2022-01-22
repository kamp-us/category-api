package helper

import (
	"github.com/kamp-us/category-api/internal/models"
	api "github.com/kamp-us/category-api/rpc/category-api"
)

func ConvertToCategoryModel(category *models.Category) *api.Category {
	return &api.Category{
		ID:          category.ID.String(),
		Name:        category.Name,
		Description: category.Description,
		Slug:        category.Slug,
	}
}
