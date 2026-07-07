package postgres

import (
	"context"

	"github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/entity"
)

func (r *Repository) Update(
	ctx context.Context,
	wallet *entity.Wallet,
) error {

	_, err := r.db.Exec(
		ctx,
		`
		UPDATE wallets
		SET
			balance=$1,
			updated_at=$2
		WHERE id=$3
		`,
		wallet.Balance.String(),
		wallet.UpdatedAt,
		wallet.ID,
	)

	return err
}