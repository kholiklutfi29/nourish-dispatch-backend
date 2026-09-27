CREATE TABLE drivers (
    user_id UUID PRIMARY KEY,

    vehicle_type VARCHAR(30) NOT NULL,

    vehicle_plate VARCHAR(20) NOT NULL UNIQUE,

    status VARCHAR(30) NOT NULL DEFAULT 'OFFLINE',

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_drivers_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT drivers_status_check
        CHECK (
            status IN (
                'OFFLINE',
                'AVAILABLE',
                'BUSY',
                'SUSPENDED'
            )
        )
);