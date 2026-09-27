CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    customer_id UUID NOT NULL,

    restaurant_id UUID NOT NULL,

    driver_id UUID,

    status VARCHAR(30) NOT NULL DEFAULT 'PENDING_PAYMENT',

    subtotal NUMERIC(12, 2) NOT NULL,

    delivery_fee NUMERIC(12, 2) NOT NULL DEFAULT 0,

    total_price NUMERIC(12, 2) NOT NULL,

    delivery_address TEXT NOT NULL,

    delivery_latitude NUMERIC(10, 7),

    delivery_longitude NUMERIC(10, 7),

    notes TEXT,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_orders_customer
        FOREIGN KEY (customer_id)
        REFERENCES customers(user_id)
        ON DELETE RESTRICT,

    CONSTRAINT fk_orders_restaurant
        FOREIGN KEY (restaurant_id)
        REFERENCES restaurants(id)
        ON DELETE RESTRICT,

    CONSTRAINT fk_orders_driver
        FOREIGN KEY (driver_id)
        REFERENCES drivers(user_id)
        ON DELETE SET NULL,

    CONSTRAINT orders_status_check
        CHECK (
            status IN (
                'PENDING_PAYMENT',
                'PAID',
                'RESTAURANT_ACCEPTED',
                'PREPARING',
                'READY_FOR_PICKUP',
                'DRIVER_ASSIGNED',
                'PICKED_UP',
                'ON_THE_WAY',
                'DELIVERED',
                'CANCELLED'
            )
        ),

    CONSTRAINT orders_subtotal_check
        CHECK (subtotal >= 0),

    CONSTRAINT orders_delivery_fee_check
        CHECK (delivery_fee >= 0),

    CONSTRAINT orders_total_price_check
        CHECK (total_price >= 0)
);