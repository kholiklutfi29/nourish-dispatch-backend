package models

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	UserID         uuid.UUID `json:"user_id" db:"user_id"`
	DefaultAddress *string   `json:"default_address" db:"default_address"`
	Latitude       *float64  `json:"latitude" db:"latitude"`
	Longitude      *float64  `json:"longitude" db:"longitude"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// using pointer (*) to accommodate null values