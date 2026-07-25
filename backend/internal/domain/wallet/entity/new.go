package entity

import (
    "time"

    "github.com/google/uuid"
    "github.com/shopspring/decimal"
)

func New(userID string) *Wallet {
    now := time.Now().UTC()

    return &Wallet{
        ID:        uuid.NewString(),
        UserID:    userID,
        Balance:   decimal.Zero,
        Currency:  "PI",
        CreatedAt: now,
        UpdatedAt: now,
    }
}