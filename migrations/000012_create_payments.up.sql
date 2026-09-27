CREATE TABLE payments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    order_id UUID NOT NULL UNIQUE,

    payment_method VARCHAR(30) NOT NULL,

    status VARCHAR(30) NOT NULL DEFAULT 'PENDING',

    amount NUMERIC(12, 2) NOT NULL,

    transaction_ref VARCHAR(100),

    paid_at TIMESTAMPTZ,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_payments_order
        FOREIGN KEY (order_id)
        REFERENCES orders(id)
        ON DELETE CASCADE,

    CONSTRAINT payments_status_check
        CHECK (
            status IN (
                'PENDING',
                'SUCCESS',
                'FAILED',
                'EXPIRED'
            )
        ),

    CONSTRAINT payments_amount_check
        CHECK (amount >= 0)
);