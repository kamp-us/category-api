package server

import (
	"context"
	api "github.com/kamp-us/category-api/rpc/category-api"
	"github.com/kamp-us/category-api/server/helper"
	"github.com/twitchtv/twirp"
)

func (s *CategoryAPIServer) GetCategoryBySlug(ctx context.Context, req *api.GetCategoryBySlugRequest) (*api.GetCategoryBySlugResponse, error) {
	if err := validateGetCategoryBySlugRequest(req); err != nil {
		return nil, err
	}

	category, err := s.backend.GetCategoryBySlug(ctx, req.Slug)
	if err != nil {
		return nil, err
	}

	return &api.GetCategoryBySlugResponse{Category: helper.ConvertToCategoryModel(category)}, nil
}

func validateGetCategoryBySlugRequest(req *api.GetCategoryBySlugRequest) error {
	if req.Slug == "" {
		return twirp.RequiredArgumentError("id")
	}
	return nil
}
