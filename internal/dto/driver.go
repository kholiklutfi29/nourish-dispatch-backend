package dto

import (
	"time"

	"github.com/google/uuid"
)

type DriverResponse struct {
	UserID       uuid.UUID `db:"user_id" json:"user_id"`
	VehicleType  string    `db:"vehicle_type" json:"vehicle_type"`
	VehiclePlate string    `db:"vehicle_plate" json:"vehicle_plate"`
	Status       string    `db:"status" json:"status"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

type UpdateDriverProfileRequest struct {
	VehicleType  string `db:"vehicle_type" json:"vehicle_type"`
	VehiclePlate string `db:"vehicle_plate" json:"vehicle_plate"`
}

type UpdateDriverStatusRequest struct {
	Status string `json:"status"`
}

type DriverDetailResponse struct {
    UserID       uuid.UUID `json:"user_id"`
    Name         string    `json:"name"`
    Email        string    `json:"email"`
    Phone        *string   `json:"phone"`
    IsActive     bool      `json:"is_active"`

    VehicleType  string    `json:"vehicle_type"`
    VehiclePlate string    `json:"vehicle_plate"`
    Status       string    `json:"status"`

    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}