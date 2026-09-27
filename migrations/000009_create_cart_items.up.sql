CREATE TABLE cart_items (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    cart_id UUID NOT NULL,

    menu_id UUID NOT NULL,

    quantity INTEGER NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_cart_items_cart
        FOREIGN KEY (cart_id)
        REFERENCES carts(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_cart_items_menu
        FOREIGN KEY (menu_id)
        REFERENCES menus(id)
        ON DELETE RESTRICT,

    CONSTRAINT cart_items_quantity_check
        CHECK (quantity > 0),

    CONSTRAINT cart_items_cart_menu_unique
        UNIQUE (cart_id, menu_id)
);