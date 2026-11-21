package dto

import (
	"time"

	"github.com/google/uuid"
)

type CustomerDetailResponse struct {
	UserID         uuid.UUID `json:"user_id" db:"user_id"`
	Name           string    `json:"name" db:"name"`
	Email          string    `json:"email" db:"email"`
	Phone          *string   `json:"phone" db:"phone"`
	Role           string  `json:"role" db:"role"`
	DefaultAddress *string   `json:"default_address" db:"default_address"`
	Latitude       *float64  `json:"latitude" db:"latitude"`
	Longitude      *float64  `json:"longitude" db:"longitude"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

type CustomerAddressResponse struct {
	UserID         uuid.UUID `json:"user_id" db:"user_id"`
	DefaultAddress *string   `json:"default_address" db:"default_address"`
	Latitude       *float64  `json:"latitude" db:"latitude"`
	Longitude      *float64  `json:"longitude" db:"longitude"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

type CustomerAddressRequest struct {
	UserID         uuid.UUID `json:"user_id" db:"user_id"`
	DefaultAddress *string   `json:"default_address" db:"default_address"`
	Latitude       *float64  `json:"latitude" db:"latitude"`
	Longitude      *float64  `json:"longitude" db:"longitude"`
}