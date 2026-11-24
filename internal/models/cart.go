package models

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID           uuid.UUID `json:"id"`
	CustomerID   uuid.UUID `json:"customer_id"`
	RestaurantID uuid.UUID `json:"restaurant_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}