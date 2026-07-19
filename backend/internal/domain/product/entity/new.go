package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// New creates a valid Product entity.
func New(
	code string,
	name string,
	category string,
	provider string,
	price decimal.Decimal,
) *Product {

	now := time.Now().UTC()

	return &Product{
		ID: uuid.New(),

		Code: code,

		Name: name,

		Category: category,

		Provider: provider,

		Price: price,

		Currency: "PI",

		Active: true,

		CreatedAt: now,
		UpdatedAt: now,
	}
}