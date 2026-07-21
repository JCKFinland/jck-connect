package handler

import (
	"github.com/JCKFinland/jck-connect/backend/internal/domain/user/service"
)

// Handler handles HTTP requests for the user domain.
type Handler struct {
	service service.Service
}

// New creates a new user handler.
func New(
	service service.Service,
) *Handler {
	return &Handler{
		service: service,
	}
}
