CREATE TABLE customers (
    user_id UUID PRIMARY KEY,

    default_address TEXT,

    latitude NUMERIC(10, 7),

    longitude NUMERIC(10, 7),

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_customers_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);