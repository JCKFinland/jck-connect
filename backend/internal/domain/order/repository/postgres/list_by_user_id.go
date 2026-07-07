package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	orderentity "github.com/JCKFinland/jck-connect/backend/internal/domain/order/entity"
)

// ListByUserID returns all orders belonging to a user,
// ordered from newest to oldest.
func (r *Repository) ListByUserID(
	ctx context.Context,
	userID uuid.UUID,
) ([]*orderentity.Order, error) {

	const query = `
	SELECT
		id,
		user_id,
		product_id,
		reference,
		amount,
		currency,
		status,
		created_at,
		updated_at
	FROM orders
	WHERE user_id = $1
	ORDER BY created_at DESC
	`

	rows, err := r.db.Query(
		ctx,
		query,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]*orderentity.Order, 0)

	for rows.Next() {

		var order orderentity.Order
		var amount string

		if err := rows.Scan(
			&order.ID,
			&order.UserID,
			&order.ProductID,
			&order.Reference,
			&amount,
			&order.Currency,
			&order.Status,
			&order.CreatedAt,
			&order.UpdatedAt,
		); err != nil {
			return nil, err
		}

		order.Amount = decimal.RequireFromString(amount)

		orders = append(orders, &order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}