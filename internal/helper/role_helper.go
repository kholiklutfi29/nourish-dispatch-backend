package helper

import "github.com/kholiklutfi29/nourish-dispatch/internal/models"

func IsValidRole(role models.UserRole) bool {

	switch role {

	case models.UserCustomer,
		models.UserMerchant,
		models.UserDriver,
		models.UserAdmin:

		return true

	default:
		return false
	}
}