CREATE TABLE products (
    id UUID PRIMARY KEY,

    code VARCHAR(100) NOT NULL UNIQUE,

    name VARCHAR(255) NOT NULL,

    category VARCHAR(50) NOT NULL,

    provider VARCHAR(100) NOT NULL,

    description TEXT,

    price NUMERIC(20,8) NOT NULL,

    currency VARCHAR(10) NOT NULL DEFAULT 'PI',

    active BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

-- Indexes

CREATE INDEX idx_products_category
    ON products(category);

CREATE INDEX idx_products_provider
    ON products(provider);

CREATE INDEX idx_products_active
    ON products(active);