package purchase

import "github.com/google/uuid"

// PurchaseRequest represents an incoming purchase request.
//
// The authenticated UserID is NOT supplied by the client.
// It is injected by the handler from the authentication context.
type PurchaseRequest struct {
	ProductID uuid.UUID `json:"product_id" binding:"required"`
}