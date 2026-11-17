package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type DriverDetail struct {
	UserID   uuid.UUID
	Name     string
	Email    string
	Phone    sql.NullString
	IsActive bool

	VehicleType  string
	VehiclePlate string
	Status       DriverStatus

	CreatedAt time.Time
	UpdatedAt time.Time
}