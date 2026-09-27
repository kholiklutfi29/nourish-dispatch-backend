CREATE TABLE driver_locations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    driver_id UUID NOT NULL,

    order_id UUID,

    latitude NUMERIC(10, 7) NOT NULL,

    longitude NUMERIC(10, 7) NOT NULL,

    recorded_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_driver_locations_driver
        FOREIGN KEY (driver_id)
        REFERENCES drivers(user_id)
        ON DELETE CASCADE,

    CONSTRAINT fk_driver_locations_order
        FOREIGN KEY (order_id)
        REFERENCES orders(id)
        ON DELETE SET NULL,

    CONSTRAINT driver_locations_latitude_check
        CHECK (latitude BETWEEN -90 AND 90),

    CONSTRAINT driver_locations_longitude_check
        CHECK (longitude BETWEEN -180 AND 180)
);