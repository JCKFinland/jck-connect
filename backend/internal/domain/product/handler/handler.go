package handler

import (
	productservice "github.com/JCKFinland/jck-connect/backend/internal/domain/product/service"
)

// Handler handles Product HTTP requests.
type Handler struct {
	service productservice.Service
}

// New creates a new Product handler.
func New(
	service productservice.Service,
) *Handler {
	return &Handler{
		service: service,
	}
}