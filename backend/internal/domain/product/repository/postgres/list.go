package postgres

import (
	"context"

	productentity "github.com/JCKFinland/jck-connect/backend/internal/domain/product/entity"
)

// List returns all products ordered by name.
func (r repository) List(
	ctx context.Context,
) ([]*productentity.Product, error) {

	const query = `
	SELECT
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
	FROM products
	ORDER BY name ASC
	`

	rows, err := r.db.Query(
		ctx,
		query,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]*productentity.Product, 0)

	for rows.Next() {

		var product productentity.Product

		if err := rows.Scan(
			&product.ID,
			&product.Code,
			&product.Name,
			&product.Category,
			&product.Provider,
			&product.Price,
			&product.Currency,
			&product.Active,
			&product.CreatedAt,
			&product.UpdatedAt,
		); err != nil {
			return nil, err
		}

		products = append(
			products,
			&product,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
