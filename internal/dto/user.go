package dto

import "github.com/kholiklutfi29/nourish-dispatch/internal/models"

type UserProfileRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserProfileResponse struct {
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Email    string          `json:"email"`
	Phone    *string         `json:"phone"`
	Role     models.UserRole `json:"role"`
	IsActive bool            `json:"is_active"`
}

type UserChangeNameRequest struct {
	NewName string `json:"new_name" binding:"required,min=2"`
}

type UserChangePhoneRequest struct {
	NewPhone    *string         `json:"new_phone"`
}
