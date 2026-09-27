CREATE TABLE menus (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    restaurant_id UUID NOT NULL,

    name VARCHAR(150) NOT NULL,

    description TEXT,

    price NUMERIC(12, 2) NOT NULL,

    image_url TEXT,

    is_available BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_menus_restaurant
        FOREIGN KEY (restaurant_id)
        REFERENCES restaurants(id)
        ON DELETE CASCADE,

    CONSTRAINT menus_price_check
        CHECK (price >= 0)
);