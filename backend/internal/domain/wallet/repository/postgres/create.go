package postgres

import (
	"context"

	"github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/entity"
)

func (r *repository) Create(
	ctx context.Context,
	wallet *entity.Wallet,
) error {

	_, err := r.db.Exec(
		ctx,
		`
		INSERT INTO wallets
		(
			id,
			user_id,
			balance,
			currency,
			created_at,
			updated_at
		)
		VALUES ($1,$2,$3,$4,$5,$6)
		`,
		wallet.ID,
		wallet.UserID,
		wallet.Balance.String(),
		wallet.Currency,
		wallet.CreatedAt,
		wallet.UpdatedAt,
	)

	return err
}