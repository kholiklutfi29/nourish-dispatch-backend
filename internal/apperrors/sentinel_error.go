package apperrors

import "errors"

var (
	ErrEmailAlreadyRegistered = errors.New("email already registered")
	ErrInvalidCredentials     = errors.New("invalid email or password")
	ErrUserNotFound           = errors.New("user not found")
	ErrAccountInactive        = errors.New("account is inactive")

	ErrInvalidRole = errors.New("invalid user role")
)