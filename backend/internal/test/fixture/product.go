package fixture

import (
	"testing"

	"github.com/shopspring/decimal"

	productentity "github.com/JCKFinland/jck-connect/backend/internal/domain/product/entity"
)

func Product(
	t *testing.T,
) *productentity.Product {

	t.Helper()

	return productentity.New(
		"TEST-PRODUCT",
		"Integration Product",
		"TEST",
		"JCK",
		decimal.NewFromInt(10),
	)
}