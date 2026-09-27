package service

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/kholiklutfi29/nourish-dispatch/internal/apperrors"
	"github.com/kholiklutfi29/nourish-dispatch/internal/models"
	"github.com/kholiklutfi29/nourish-dispatch/internal/repository"
)


type ProfileService struct {
	customerRepository *repository.CustomerRepository
	driverRepository     *repository.DriverRepository
	restaurantRepository *repository.RestaurantRepository
}

func NewProfileService(
	customerRepository *repository.CustomerRepository,
	driverRepository *repository.DriverRepository,
	restaurantRepository *repository.RestaurantRepository,
) *ProfileService {
	return &ProfileService{
		driverRepository: driverRepository,
		restaurantRepository: restaurantRepository,
	}
}

func (s *ProfileService) CreateForRoleTx(
	ctx context.Context,
	tx *sql.Tx,
	user *models.User,
) error {

	switch user.Role {

	case models.UserCustomer:
		return s.createCustomerTx(ctx, tx, user)

	case models.UserMerchant:
		return s.createMerchantTx(ctx, tx, user)

	case models.UserDriver:
		return s.createDriverTx(ctx, tx, user)

	case models.UserAdmin:
		// Admin tidak membutuhkan
		// tabel profile tambahan.
		return nil

	default:
		return apperrors.ErrInvalidRole
	}
}


func (s *ProfileService) createMerchantTx(
	ctx context.Context,
	tx *sql.Tx,
	user *models.User,
) error {

	restaurant := &models.Restaurant{
		ID:       uuid.New(),
		OwnerID:  user.ID,
		Name:     "",
		Address:  "",
		IsOpen:   false,
		IsActive: true,
	}

	return s.restaurantRepository.CreateTx(
		ctx,
		tx,
		restaurant,
	)
}

func (s *ProfileService) createCustomerTx(
	ctx context.Context,
	tx *sql.Tx,
	user *models.User,
) error {

	customer := &models.Customer{
		UserID: user.ID,
		DefaultAddress: nil,
		Latitude: nil,
		Longitude: nil,
	}

	return s.customerRepository.CreateTx(
		ctx,
		tx, 
		customer,
	)

}

func (s *ProfileService) createDriverTx(
	ctx context.Context,
	tx *sql.Tx,
	user *models.User,
) error {

	driver := &models.Driver{
		UserID:       user.ID,
		VehicleType:  "",
		VehiclePlate: "",
		Status:       models.DriverStatusOffline,
	}

	return s.driverRepository.CreateTx(
		ctx,
		tx,
		driver,
	)
}