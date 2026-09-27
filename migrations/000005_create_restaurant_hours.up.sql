CREATE TABLE restaurant_hours (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    restaurant_id UUID NOT NULL,

    day_of_week SMALLINT NOT NULL,

    open_time TIME,

    close_time TIME,

    is_closed BOOLEAN NOT NULL DEFAULT FALSE,

    CONSTRAINT fk_restaurant_hours_restaurant
        FOREIGN KEY (restaurant_id)
        REFERENCES restaurants(id)
        ON DELETE CASCADE,

    CONSTRAINT restaurant_hours_day_check
        CHECK (day_of_week BETWEEN 0 AND 6),

    CONSTRAINT restaurant_hours_time_check
        CHECK (
            is_closed = TRUE
            OR (
                open_time IS NOT NULL
                AND close_time IS NOT NULL
            )
        )
);