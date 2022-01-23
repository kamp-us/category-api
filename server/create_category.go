package server

import (
	"context"
	api "github.com/kamp-us/category-api/rpc/category-api"
	"github.com/kamp-us/category-api/server/helper"
	"github.com/twitchtv/twirp"
)

func (s *CategoryAPIServer) CreateCategory(ctx context.Context, req *api.CreateCategoryRequest) (*api.Category, error) {
	if err := validateCreateCategoryRequest(req); err != nil {
		return nil, err
	}

	category, err := s.backend.CreateCategory(ctx, req.Name, req.Description)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return helper.ConvertToCategoryModel(category), nil
}

func validateCreateCategoryRequest(req *api.CreateCategoryRequest) error {
	if req.Name == "" {
		return twirp.RequiredArgumentError("name")
	}
	if req.Description == "" {
		return twirp.RequiredArgumentError("description")
	}
	return nil
}
