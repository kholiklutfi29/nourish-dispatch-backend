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

	ErrDriverNotFound = errors.New("driver not found")
	ErrInvalidDriverStatus = errors.New("driver status not valid")

	ErrCustomerNotFound = errors.New("customer not found")

	ErrMenuNotFound = errors.New("menu not found")

	ErrInvalidMenuName = errors.New("menu name is required")

	ErrInvalidMenuPrice = errors.New("menu price is required")

	ErrCartNotFound = errors.New("cart not found")

	ErrCartItemNotFound = errors.New("cart item not found")
)