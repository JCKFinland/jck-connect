package postgres

import (
	"context"

	orderentity "github.com/JCKFinland/jck-connect/backend/internal/domain/order/entity"
)

// Create inserts a new order into the database.
func (r *Repository) Create(
	ctx context.Context,
	order *orderentity.Order,
) error {

	const query = `
	INSERT INTO orders (
		id,
		user_id,
		product_id,
		reference,
		amount,
		currency,
		status,
		created_at,
		updated_at
	)
	VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9
	)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		order.ID,
		order.UserID,
		order.ProductID,
		order.Reference,
		order.Amount.String(),
		order.Currency,
		order.Status,
		order.CreatedAt,
		order.UpdatedAt,
	)

	return err
}