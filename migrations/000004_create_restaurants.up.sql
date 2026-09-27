CREATE TABLE restaurants (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    owner_id UUID NOT NULL,

    name VARCHAR(150) NOT NULL,

    description TEXT,

    phone VARCHAR(30),

    address TEXT NOT NULL,

    latitude NUMERIC(10, 7),

    longitude NUMERIC(10, 7),

    is_open BOOLEAN NOT NULL DEFAULT FALSE,

    is_active BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_restaurants_owner
        FOREIGN KEY (owner_id)
        REFERENCES users(id)
        ON DELETE RESTRICT
);