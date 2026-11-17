package apperrors

import "errors"

var (
	ErrEmailAlreadyRegistered = errors.New("email already registered")
	ErrInvalidCredentials     = errors.New("invalid email or password")
	ErrUserNotFound           = errors.New("user not found")
	ErrAccountInactive        = errors.New("account is inactive")

	ErrInvalidRole = errors.New("invalid user role")

	ErrInternal = errors.New("internal server error")

	ErrRestaurantNotFound = errors.New("restaurant not found")
	ErrInvalidRestaurantName = errors.New("restaurant name is required")
	ErrInvalidRestaurantOwnerID = errors.New("invalid restaurant owner id")
	ErrInvalidRestaurantData = errors.New("invalid restaurant data")
)