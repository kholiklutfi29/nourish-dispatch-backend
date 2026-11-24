package models

import (
	"time"

	"github.com/google/uuid"
)

type Menu struct {
	ID           uuid.UUID `json:"id" db:"id"`
	RestaurantID uuid.UUID `json:"restaurant_id" db:"restaurant_id"`
	Name         string    `json:"name" db:"name"`
	Description  *string   `json:"description" db:"description"`
	Price        float64   `json:"price" db:"price"`
	ImageURL     *string   `json:"image_url" db:"image_url"`
	IsAvailable  bool      `json:"is_available" db:"is_available"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at" db:"deleted_at"`
}