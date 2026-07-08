package handler

import (
	orderservice "github.com/JCKFinland/jck-connect/backend/internal/domain/order/service"
)

// Handler handles HTTP requests for the Order domain.
type Handler struct {
	service orderservice.Service
}

// New creates a new Order handler.
func New(
	service orderservice.Service,
) *Handler {
	return &Handler{
		service: service,
	}
}
