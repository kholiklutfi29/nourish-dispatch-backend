CREATE TABLE order_status_history (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    order_id UUID NOT NULL,

    changed_by UUID NOT NULL,

    status VARCHAR(30) NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_order_status_history_order
        FOREIGN KEY (order_id)
        REFERENCES orders(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_order_status_history_user
        FOREIGN KEY (changed_by)
        REFERENCES users(id)
        ON DELETE RESTRICT,

    CONSTRAINT order_status_history_status_check
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
        )
);