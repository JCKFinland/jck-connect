CREATE TABLE orders (
    id UUID PRIMARY KEY,

    user_id UUID NOT NULL,

    product_id UUID NOT NULL,

    amount NUMERIC(20,8) NOT NULL,

    currency VARCHAR(10) NOT NULL DEFAULT 'PI',

    status VARCHAR(30) NOT NULL,

    reference VARCHAR(100) NOT NULL UNIQUE,

    provider_reference VARCHAR(100),

    metadata JSONB,

    created_at TIMESTAMPTZ NOT NULL,

    updated_at TIMESTAMPTZ NOT NULL,

    CONSTRAINT fk_orders_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE NO ACTION,

    CONSTRAINT fk_orders_product
        FOREIGN KEY (product_id)
        REFERENCES products(id)
        ON DELETE NO ACTION,

    CONSTRAINT chk_orders_status
        CHECK (
            status IN (
                'PENDING',
                'PROCESSING',
                'COMPLETED',
                'FAILED',
                'CANCELLED'
            )
        )
);

CREATE INDEX idx_orders_user_id
    ON orders(user_id);

CREATE INDEX idx_orders_product_id
    ON orders(product_id);

CREATE INDEX idx_orders_status
    ON orders(status);

CREATE INDEX idx_orders_created_at
    ON orders(created_at);