CREATE TABLE wallets (
    id UUID PRIMARY KEY,

    user_id UUID NOT NULL UNIQUE,

    balance NUMERIC(20,8) NOT NULL DEFAULT 0,

    currency VARCHAR(10) NOT NULL DEFAULT 'PI',

    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,

    CONSTRAINT fk_wallets_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

-- Indexes

CREATE INDEX idx_wallets_currency
    ON wallets(currency);