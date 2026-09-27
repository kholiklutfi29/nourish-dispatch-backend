CREATE TABLE order_items (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    order_id UUID NOT NULL,

    menu_id UUID NOT NULL,

    menu_name VARCHAR(150) NOT NULL,

    unit_price NUMERIC(12, 2) NOT NULL,

    quantity INTEGER NOT NULL,

    subtotal NUMERIC(12, 2) NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_order_items_order
        FOREIGN KEY (order_id)
        REFERENCES orders(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_order_items_menu
        FOREIGN KEY (menu_id)
        REFERENCES menus(id)
        ON DELETE RESTRICT,

    CONSTRAINT order_items_unit_price_check
        CHECK (unit_price >= 0),

    CONSTRAINT order_items_quantity_check
        CHECK (quantity > 0),

    CONSTRAINT order_items_subtotal_check
        CHECK (subtotal >= 0)
);