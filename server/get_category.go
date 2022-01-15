package server

import (
	"context"
	api "github.com/kamp-us/category-api/rpc/category-api"
	"github.com/twitchtv/twirp"
)

func (s *CategoryAPIServer) GetCategory(ctx context.Context, req *api.GetCategoryRequest) (*api.Category, error) {

	if err := validateGetCategoryRequest(req); err != nil {
		return nil, err
	}

	Category, err := s.backend.GetCategory(ctx, req.ID)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &api.Category{
		ID:          Category.ID.String(),
		Name:        Category.Name,
		Description: Category.Description,
		Slug:        Category.Slug,
	}, nil

}

func validateGetCategoryRequest(req *api.GetCategoryRequest) error {
	if req.ID == "" {
		return twirp.RequiredArgumentError("id")
	}
	return nil
}
