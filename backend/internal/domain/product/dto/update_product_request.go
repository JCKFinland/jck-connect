package dto

import "github.com/shopspring/decimal"

// UpdateProductRequest represents a request to update a product.
type UpdateProductRequest struct {
	Code     string          `json:"code" binding:"required"`
	Name     string          `json:"name" binding:"required"`
	Category string          `json:"category" binding:"required"`
	Provider string          `json:"provider" binding:"required"`
	Price    decimal.Decimal `json:"price" binding:"required"`
	Currency string          `json:"currency" binding:"required"`
	Active   bool            `json:"active"`
}
