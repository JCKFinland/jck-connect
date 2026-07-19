package fixture

import (
	"context"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"

	productentity "github.com/JCKFinland/jck-connect/backend/internal/domain/product/entity"
	productrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/product/repository"
)

func CreateProduct(
	t *testing.T,
	repo productrepo.Repository,
) *productentity.Product {

	t.Helper()

	product := productentity.New(
		"TEST-PRODUCT",
		"Integration Product",
		"TEST",
		"JCK",
		decimal.NewFromInt(10),
	)

	err := repo.Create(
		context.Background(),
		product,
	)

	require.NoError(t, err)

	return product
}