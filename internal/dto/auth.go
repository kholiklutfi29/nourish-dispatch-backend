package dto

import "github.com/kholiklutfi29/nourish-dispatch/internal/models"

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Phone    string `json:"phone"`

	// lock just 3 role, excluding admin for security
	Role models.UserRole `json:"role" binding:"required,oneof=CUSTOMER MERCHANT DRIVER"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Phone    *string `json:"phone"`
	Role     models.UserRole  `json:"role"`
	IsActive bool    `json:"is_active"`
}

type AuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type RegisterResponse struct {
	User UserResponse `json:"user"`
}