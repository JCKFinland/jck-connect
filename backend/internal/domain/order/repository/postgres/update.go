package postgres

import (
	"context"

	orderentity "github.com/JCKFinland/jck-connect/backend/internal/domain/order/entity"
)

// Update persists changes to an existing order.
func (r *Repository) Update(
	ctx context.Context,
	order *orderentity.Order,
) error {

	const query = `
	UPDATE orders
	SET
		reference = $1,
		amount = $2,
		currency = $3,
		status = $4,
		updated_at = $5
	WHERE id = $6
	`

	_, err := r.db.Exec(
		ctx,
		query,
		order.Reference,
		order.Amount.String(),
		order.Currency,
		order.Status,
		order.UpdatedAt,
		order.ID,
	)

	return err
}
