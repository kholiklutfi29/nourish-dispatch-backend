package models

import (
	"time"

	"github.com/google/uuid"
)

type DriverStatus string

const (
	DriverStatusOffline   DriverStatus = "OFFLINE"
	DriverStatusAvailable DriverStatus = "AVAILABLE"
	DriverStatusBusy      DriverStatus = "BUSY"
	DriverStatusSuspended DriverStatus = "SUSPENDED"
)

type Driver struct {
	UserID       uuid.UUID    `db:"user_id" json:"user_id"`
	VehicleType  string       `db:"vehicle_type" json:"vehicle_type"`
	VehiclePlate string       `db:"vehicle_plate" json:"vehicle_plate"`
	Status       DriverStatus `db:"status" json:"status"`
	CreatedAt    time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time    `db:"updated_at" json:"updated_at"`
}