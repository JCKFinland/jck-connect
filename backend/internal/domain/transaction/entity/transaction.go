package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// TransactionType represents the type of transaction.
type TransactionType string

const (
	TransactionTypeDebit  TransactionType = "DEBIT"
	TransactionTypeCredit TransactionType = "CREDIT"
	TransactionTypeRefund TransactionType = "REFUND"
)

// TransactionStatus represents the transaction status.
type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "PENDING"
	TransactionStatusCompleted TransactionStatus = "COMPLETED"
	TransactionStatusFailed    TransactionStatus = "FAILED"
)

// Transaction represents a wallet ledger transaction.
//
// Transactions are immutable financial records.
// Existing transactions should never be modified.
// Corrections must be represented by new transactions.
type Transaction struct {
	ID uuid.UUID

	WalletID string

	OrderID uuid.UUID

	Reference string

	BalanceBefore decimal.Decimal

	BalanceAfter decimal.Decimal

	Type TransactionType

	Amount decimal.Decimal

	Currency string

	Description string

	Status TransactionStatus

	CreatedAt time.Time
}
