package server

import (
	"github.com/kamp-us/category-api/internal/backend"
)

type CategoryAPIServer struct {
	backend backend.Backender
}

func NewCategoryAPIServer(backend backend.Backender) *CategoryAPIServer {
	return &CategoryAPIServer{backend: backend}
}
