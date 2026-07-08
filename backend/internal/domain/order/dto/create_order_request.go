package dto

import (
	"github.com/google/uuid"
)

type CreateOrderRequest struct {
	ProductID uuid.UUID `json:"product_id" binding:"required"`
}
