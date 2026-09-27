CREATE TABLE driver_assignments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    order_id UUID NOT NULL,

    driver_id UUID NOT NULL,

    status VARCHAR(30) NOT NULL,

    assigned_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    accepted_at TIMESTAMPTZ,

    rejected_at TIMESTAMPTZ,

    completed_at TIMESTAMPTZ,

    CONSTRAINT fk_driver_assignments_order
        FOREIGN KEY (order_id)
        REFERENCES orders(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_driver_assignments_driver
        FOREIGN KEY (driver_id)
        REFERENCES drivers(user_id)
        ON DELETE RESTRICT,

    CONSTRAINT driver_assignments_status_check
        CHECK (
            status IN (
                'OFFERED',
                'ACCEPTED',
                'REJECTED',
                'COMPLETED'
            )
        )
);