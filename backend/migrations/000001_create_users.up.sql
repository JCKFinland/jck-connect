CREATE TABLE users (
    id UUID PRIMARY KEY,

    pi_uid VARCHAR(100) NOT NULL UNIQUE,
    pi_username VARCHAR(100) NOT NULL UNIQUE,

    display_name VARCHAR(100) NOT NULL,

    email VARCHAR(255),
    phone_number VARCHAR(30),

    role VARCHAR(30) NOT NULL,
    status VARCHAR(30) NOT NULL,

    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

-- Additional indexes

CREATE INDEX idx_users_email
    ON users(email);

CREATE INDEX idx_users_status
    ON users(status);
    