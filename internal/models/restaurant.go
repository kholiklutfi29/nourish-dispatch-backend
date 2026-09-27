package models

import (
	"time"

	"github.com/google/uuid"
)

type Restaurant struct {
	ID          uuid.UUID `db:"id" json:"id"`
	OwnerID     uuid.UUID `db:"owner_id" json:"owner_id"`
	Name        string    `db:"name" json:"name"`
	Description *string   `db:"description" json:"description,omitempty"`
	Phone       *string   `db:"phone" json:"phone,omitempty"`
	Address     string    `db:"address" json:"address"`
	Latitude    *float64  `db:"latitude" json:"latitude,omitempty"`
	Longitude   *float64  `db:"longitude" json:"longitude,omitempty"`
	IsOpen      bool      `db:"is_open" json:"is_open"`
	IsActive    bool      `db:"is_active" json:"is_active"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}