package models

import (
	"time"

	"github.com/google/uuid"
)

type UserRole string

const (
	UserCustomer UserRole = "CUSTOMER"
	UserMerchant UserRole = "MERCHANT"
	UserDriver   UserRole = "DRIVER"
	UserAdmin    UserRole = "ADMIN"
)

type User struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"-" db:"password_hash"`
	Phone        *string   `json:"phone" db:"phone"`
	Role         UserRole    `json:"role" db:"role"`
	IsActive     bool      `json:"is_active" db:"is_active"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
