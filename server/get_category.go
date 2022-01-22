package server

import (
	"context"
	api "github.com/kamp-us/category-api/rpc/category-api"
	"github.com/kamp-us/category-api/server/helper"
	"github.com/twitchtv/twirp"
)

func (s *CategoryAPIServer) GetCategory(ctx context.Context, req *api.GetCategoryRequest) (*api.Category, error) {

	if err := validateGetCategoryRequest(req); err != nil {
		return nil, err
	}

	category, err := s.backend.GetCategory(ctx, req.ID)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return helper.ConvertToCategoryModel(category), nil

}

func validateGetCategoryRequest(req *api.GetCategoryRequest) error {
	if req.ID == "" {
		return twirp.RequiredArgumentError("id")
	}
	return nil
}
