package purchase

import (
	purchaseusecase "github.com/JCKFinland/jck-connect/backend/internal/usecase/purchase"
)

// Handler exposes HTTP endpoints for the Purchase use case.
type Handler struct {
	purchaseService purchaseusecase.Service
}

// New creates a new Purchase HTTP handler.
func New(
	purchaseService purchaseusecase.Service,
) *Handler {

	return &Handler{
		purchaseService: purchaseService,
	}
}