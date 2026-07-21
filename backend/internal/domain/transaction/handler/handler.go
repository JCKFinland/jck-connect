package handler

import (
	transactionservice "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/service"
)

// Handler handles transaction HTTP requests.
type Handler struct {
	service transactionservice.Service
}

// New creates a new transaction handler.
func New(
	service transactionservice.Service,
) *Handler {

	return &Handler{
		service: service,
	}
}
