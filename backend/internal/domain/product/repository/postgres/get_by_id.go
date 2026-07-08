package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	productentity "github.com/JCKFinland/jck-connect/backend/internal/domain/product/entity"
	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"
)

// GetByID returns a product by its ID.
func (r *Repository) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*productentity.Product, error) {

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
	WHERE id = $1
	`

	var product productentity.Product

	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
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
	)

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, sharedErrors.New(
				sharedErrors.CodeNotFound,
				"Product not found.",
				sharedErrors.ErrNotFound,
			)
		}

		return nil, err
	}

	return &product, nil
}