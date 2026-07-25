package postgres

import (
	walletrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/repository"
	"github.com/JCKFinland/jck-connect/backend/pkg/database"
)

type repository struct {
	db database.DBTX
}

var _ walletrepo.Repository = (*repository)(nil)

func New(db *database.Database) walletrepo.Repository {
	return &repository{
		db: db,
	}
}

func NewTx(tx database.DBTX) walletrepo.Repository {
	return &repository{
		db: tx,
	}
}