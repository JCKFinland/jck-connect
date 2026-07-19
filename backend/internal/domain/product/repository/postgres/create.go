package postgres

import (
	"context"

	productentity "github.com/JCKFinland/jck-connect/backend/internal/domain/product/entity"
)

// Create inserts a new product.
func (r repository) Create(
	ctx context.Context,
	product *productentity.Product,
) error {

	const query = `
	INSERT INTO products
	(
		id,
		code,
		name,
		category,
		provider,
		price,
		currency,
		active,
		created_at,
		updated_at
	)
	VALUES
	(
		$1,$2,$3,$4,$5,$6,$7,$8,$9,$10
	)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		product.ID,
		product.Code,
		product.Name,
		product.Category,
		product.Provider,
		product.Price.String(),
		product.Currency,
		product.Active,
		product.CreatedAt,
		product.UpdatedAt,
	)

	return err
}