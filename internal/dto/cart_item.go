package dto

import (
	"time"

	"github.com/google/uuid"
)

type CartItemMenuResponse struct {
	ID        uuid.UUID `json:"id"`
	CartID    uuid.UUID `json:"cart_id"`
	MenuID    uuid.UUID `json:"menu_id"`
	MenuName  string    `json:"menu_name"`
	MenuImage string    `json:"menu_image"`
	Price     float64   `json:"price"`
	Quantity  int       `json:"quantity"`
	Subtotal  float64   `json:"subtotal"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CartItemResponse struct {
	ID        uuid.UUID `json:"id"`
	CartID    uuid.UUID `json:"cart_id"`
	MenuID    uuid.UUID `json:"menu_id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateCartItemRequest struct {
	MenuID   string `json:"menu_id" binding:"required"`
	Quantity int    `json:"quantity" binding:"required,min=1"`
}

type UpdateCartItemRequest struct {
	Quantity int `json:"quantity" binding:"required,min=1"`
}

type AddCartItemRequest struct {
	MenuID   uuid.UUID `json:"menu_id" binding:"required"`
	Quantity int    `json:"quantity" binding:"required,min=1"`
}
