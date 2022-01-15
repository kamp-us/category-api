package server

import (
	"context"
	api "github.com/kamp-us/category-api/rpc/category-api"
	"github.com/twitchtv/twirp"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (s *CategoryAPIServer) DeleteCategory(ctx context.Context, req *api.DeleteCategoryRequest) (*emptypb.Empty, error) {
	if err := validateDeleteCategoryRequest(req); err != nil {
		return nil, err
	}
	err := s.backend.DeleteCategory(ctx, req.ID)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}
	return &emptypb.Empty{}, nil
}

func convertToStringPtr(value *wrapperspb.StringValue) *string {
	val := value.GetValue()
	return &val
}

func validateDeleteCategoryRequest(req *api.DeleteCategoryRequest) error {
	if req.ID == "" {
		return twirp.RequiredArgumentError("id")
	}
	return nil
}
