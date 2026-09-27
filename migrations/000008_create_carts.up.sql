CREATE TABLE carts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    customer_id UUID NOT NULL,

    restaurant_id UUID NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_carts_customer
        FOREIGN KEY (customer_id)
        REFERENCES customers(user_id)
        ON DELETE CASCADE,

    CONSTRAINT fk_carts_restaurant
        FOREIGN KEY (restaurant_id)
        REFERENCES restaurants(id)
        ON DELETE RESTRICT,

    CONSTRAINT carts_customer_unique
        UNIQUE (customer_id)
);