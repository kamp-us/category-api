package server

import (
	"context"
	api "github.com/kamp-us/category-api/rpc/category-api"
	"github.com/twitchtv/twirp"
)

func (s CategoryAPIServer) GetBatchCategories(ctx context.Context, req *api.GetBatchCategoriesRequest) (*api.GetBatchCategoriesResponse, error) {
	if err := validateGetBatchCategoriesRequest(req); err != nil {
		return nil, err
	}
	categories, err := s.backend.GetBatchCategories(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	var batch []*api.Category
	for _, model := range categories {
		category := &api.Category{
			ID:          model.ID.String(),
			Name:        model.Name,
			Description: model.Description,
			Slug:        model.Slug,
		}
		batch = append(batch, category)
	}
	return &api.GetBatchCategoriesResponse{Categories: batch}, nil

}

func validateGetBatchCategoriesRequest(req *api.GetBatchCategoriesRequest) error {
	if len(req.Ids) == 0 {
		return twirp.RequiredArgumentError("ids")
	}
	return nil
}
