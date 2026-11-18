package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/kholiklutfi29/nourish-dispatch/internal/apperrors"
	"github.com/kholiklutfi29/nourish-dispatch/internal/dto"
	"github.com/kholiklutfi29/nourish-dispatch/internal/models"
	"github.com/kholiklutfi29/nourish-dispatch/internal/repository"
)

type DriverService struct {
	db                *sql.DB
	driverRespository *repository.DriverRepository
}

func NewDriverService(
	db *sql.DB,
	driverRepository *repository.DriverRepository,
) *DriverService {
	return &DriverService{
		db:                db,
		driverRespository: driverRepository,
	}
}

func (s *DriverService) GetDriverProfile(
	ctx context.Context,
	driverID uuid.UUID,
) (*dto.DriverResponse, error) {
	// get driver profile
	driver, err := s.driverRespository.GetDriverProfile(
		ctx, driverID,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.ErrDriverNotFound
		}
		return nil,errors.New(
			"failed to get driver profile",
		)
	}

	// map model to response dto
	return &dto.DriverResponse{
		UserID: driver.UserID,
		VehicleType: driver.VehicleType,
		VehiclePlate: driver.VehiclePlate,
		Status: string(driver.Status),
		CreatedAt: driver.CreatedAt,
		UpdatedAt: driver.UpdatedAt,
	}, nil
}

func (s *DriverService) UpdateDriverProfile(
	ctx context.Context,
	driverID uuid.UUID,
	vehicleType string,
	vehiblePlate string,
) error {
	// call repository
	err := s.driverRespository.UpdateDriverProfile(
		ctx, driverID, vehicleType, vehiblePlate,
	)

	if err != nil {
		if errors.Is(err, apperrors.ErrDriverNotFound) {
			return apperrors.ErrDriverNotFound
		}

		return errors.New("Failed to update driver profile")
	}

	return nil
}

func (s *DriverService) UpdateDriverStatus(
	ctx context.Context,
	driverID uuid.UUID,
	status models.DriverStatus,
) error {
	if status != "OFFLINE" && status != "AVAILABLE" {
		return apperrors.ErrInvalidDriverStatus
	}
	
	return s.driverRespository.UpdateDriverStatus(
		ctx, driverID, status,
	)
}

func (s *DriverService) GetAllDrivers(
	ctx context.Context,
) ([]dto.DriverResponse, error) {
	// call repository
	drivers, err := s.driverRespository.GetAllDrivers(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch drivers: %w", err)
	}

	// map model to response dto
	responses := make([]dto.DriverResponse, 0, len(drivers))
	for _, driver := range drivers {
		response:= dto.DriverResponse{
			UserID: driver.UserID,
			VehicleType: driver.VehicleType,
			VehiclePlate: driver.VehiclePlate,
			Status: string(driver.Status),
			CreatedAt: driver.CreatedAt,
			UpdatedAt: driver.UpdatedAt,
		}
		responses = append(responses, response)
	}
	return responses, nil
}

func (s *DriverService) GetDriverByStatus(
	ctx context.Context,
	status models.DriverStatus,
) ([]dto.DriverResponse, error) {
	// call repository
	drivers, err := s.driverRespository.GetDriverByStatus(
		ctx, status,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch drivers: %w", err)
	}

	// map model to response
	responses := make([]dto.DriverResponse, 0, len(drivers))
	for _, driver := range drivers {
		response := dto.DriverResponse{
			UserID: driver.UserID,
			VehicleType: driver.VehicleType,
			VehiclePlate: driver.VehiclePlate,
			Status: string(driver.Status),
			CreatedAt: driver.CreatedAt,
			UpdatedAt: driver.UpdatedAt,
		}
		responses = append(responses, response)
	}
	return responses, nil
}

func (s *DriverService) GetDriverDetails(
	ctx context.Context,
	driverID uuid.UUID,
) (*dto.DriverDetailResponse, error) {
	// call repository
	driver, err := s.driverRespository.GetDriverDetails(
		ctx, driverID,
	)

	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, apperrors.ErrDriverNotFound
        }

        return nil, fmt.Errorf(
            "failed to get driver details: %w", err,
        )
    }

	// handle nullable phone
	var phone *string
	if driver.Phone.Valid {
		phone = &driver.Phone.String
	}

	// map model to response dto
	return &dto.DriverDetailResponse{
		UserID: driver.UserID,
		Name: driver.Name,
		Email: driver.Email,
		Phone: phone,
		IsActive: driver.IsActive,
		VehicleType: driver.VehicleType,
		VehiclePlate: driver.VehiclePlate,
		Status: string(driver.Status),
		CreatedAt: driver.CreatedAt,
		UpdatedAt: driver.UpdatedAt,
	}, nil
}

