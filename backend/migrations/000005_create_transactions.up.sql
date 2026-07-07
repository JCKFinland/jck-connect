CREATE TABLE transactions (
    id UUID PRIMARY KEY,

    order_id UUID NOT NULL,

    wallet_id UUID NOT NULL,

    type VARCHAR(30) NOT NULL,

    status VARCHAR(30) NOT NULL,

    amount NUMERIC(20,8) NOT NULL,

    currency VARCHAR(10) NOT NULL DEFAULT 'PI',

    balance_before NUMERIC(20,8) NOT NULL,

    balance_after NUMERIC(20,8) NOT NULL,

    reference VARCHAR(100) NOT NULL UNIQUE,

    description TEXT,

    created_at TIMESTAMPTZ NOT NULL,

    CONSTRAINT fk_transactions_order
        FOREIGN KEY (order_id)
        REFERENCES orders(id)
        ON DELETE NO ACTION,

    CONSTRAINT fk_transactions_wallet
        FOREIGN KEY (wallet_id)
        REFERENCES wallets(id)
        ON DELETE NO ACTION,

    CONSTRAINT chk_transactions_type
        CHECK (
            type IN (
                'DEBIT',
                'CREDIT',
                'REFUND'
            )
        ),

    CONSTRAINT chk_transactions_status
        CHECK (
            status IN (
                'PENDING',
                'COMPLETED',
                'FAILED'
            )
        )
);

CREATE INDEX idx_transactions_order_id
    ON transactions(order_id);

CREATE INDEX idx_transactions_wallet_id
    ON transactions(wallet_id);

CREATE INDEX idx_transactions_reference
    ON transactions(reference);

CREATE INDEX idx_transactions_status
    ON transactions(status);

CREATE INDEX idx_transactions_created_at
    ON transactions(created_at);