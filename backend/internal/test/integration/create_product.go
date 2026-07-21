package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	productentity "github.com/JCKFinland/jck-connect/backend/internal/domain/product/entity"
	"github.com/JCKFinland/jck-connect/backend/internal/test/fixture"
)

func CreateProduct(
	t *testing.T,
	app *TestApp,
) *productentity.Product {

	t.Helper()

	product := fixture.Product(t)

	err := app.Container.ProductRepository.Create(
		context.Background(),
		product,
	)

	require.NoError(t, err)

	return product
}
