package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"

	productentity "github.com/JCKFinland/jck-connect/backend/internal/domain/product/entity"
)

// Update updates an existing product.
func (r *Repository) Update(
	ctx context.Context,
	product *productentity.Product,
) error {

	const query = `
	UPDATE products
	SET
		code = $2,
		name = $3,
		category = $4,
		provider = $5,
		price = $6,
		currency = $7,
		active = $8,
		updated_at = $9
	WHERE id = $1
	`

	commandTag, err := r.db.Exec(
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
		product.UpdatedAt,
	)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
