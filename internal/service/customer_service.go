package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/kholiklutfi29/nourish-dispatch/internal/apperrors"
	"github.com/kholiklutfi29/nourish-dispatch/internal/dto"
	"github.com/kholiklutfi29/nourish-dispatch/internal/repository"
)

type CustomerService struct {
	db                 *sql.DB
	customerRepository *repository.CustomerRepository
}

func NewCustomerService(
	db *sql.DB,
	customerRepository *repository.CustomerRepository,
) *CustomerService {
	return &CustomerService{
		db:                 db,
		customerRepository: customerRepository,
	}
}

func (s *CustomerService) GetCustomerProfile(
	ctx context.Context,
	customerID uuid.UUID,
) (*dto.CustomerDetailResponse, error) {
	// call repository
	customer, err := s.customerRepository.GetCustomerProfile(
		ctx, customerID,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.ErrCustomerNotFound
		}

		return nil, errors.New(
			"failed to get customer by id",
		)
	}

	// map model to response dto
	return &dto.CustomerDetailResponse{
		UserID:         customer.UserID,
		Name:           customer.Name,
		Email:          customer.Email,
		Phone:          customer.Phone,
		Role:           string(customer.Role),
		DefaultAddress: customer.DefaultAddress,
		Latitude:       customer.Latitude,
		Longitude:      customer.Longitude,
		CreatedAt:      customer.CreatedAt,
		UpdatedAt:      customer.UpdatedAt,
	}, nil
}

func (s *CustomerService) GetCostumerAddress(
	ctx context.Context,
	customerID uuid.UUID,
) (*dto.CustomerAddressResponse, error) {
	// call repository
	customer, err := s.customerRepository.GetCustomerAddress(
		ctx, customerID,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.ErrCustomerNotFound
		}

		return nil, errors.New(
			"failed to get customer by id",
		)
	}

	// map model to dto
	return &dto.CustomerAddressResponse{
		UserID: customerID,
		DefaultAddress: customer.DefaultAddress,
		Latitude: customer.Latitude,
		Longitude: customer.Longitude,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}, nil
}

func (s *CustomerService) UpdateCustomerAddress(
	ctx context.Context,
	customerID uuid.UUID,
	defaultAddress string,
	latitude float64,
	longitude float64,
) error {
	// call repository
	err := s.customerRepository.UpdateCustomerAddress(
		ctx, customerID, defaultAddress, latitude, longitude,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return apperrors.ErrCustomerNotFound
		}

		return errors.New(
			"failed to get customer by id",
		)
	}

	return nil
}