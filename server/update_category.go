package server

import (
	"context"

	api "github.com/kamp-us/category-api/rpc/category-api"

	"github.com/twitchtv/twirp"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *CategoryAPIServer) UpdateCategory(ctx context.Context, req *api.UpdateCategoryRequest) (*emptypb.Empty, error) {

	if err := validateUpdateCategoryRequest(req); err != nil {
		return nil, err
	}

	if err := s.backend.UpdateCategory(ctx,
		req.ID,
		convertToStringPtr(req.Name),
		convertToStringPtr(req.Description)); err != nil {
		return nil, twirp.InternalErrorWith(err)
	}
	return &emptypb.Empty{}, nil
}

func validateUpdateCategoryRequest(req *api.UpdateCategoryRequest) error {
	if req.ID == "" {
		return twirp.RequiredArgumentError("id")
	}
	return nil
}
