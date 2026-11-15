package service

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/kholiklutfi29/nourish-dispatch/internal/apperrors"
	"github.com/kholiklutfi29/nourish-dispatch/internal/dto"
	"github.com/kholiklutfi29/nourish-dispatch/internal/models"
	"github.com/kholiklutfi29/nourish-dispatch/internal/repository"
)

type RestaurantService struct {
	db                   *sql.DB
	restaurantRepository *repository.RestaurantRepository
}

func NewRestaurantService(
	db *sql.DB,
	restaurantRepository *repository.RestaurantRepository,
) *RestaurantService {
	return &RestaurantService{
		db:                   db,
		restaurantRepository: restaurantRepository,
	}
}

func (s *RestaurantService) CreateRestaurant(
	ctx context.Context,
	ownerID uuid.UUID,
	req dto.CreateRestaurantRequest,
) (*dto.RestaurantResponse, error) {
	// ==========================================
	// 1. NORMALIZE INPUT
	// ==========================================
	name := strings.TrimSpace(req.Name)
	description := ""
	if req.Description != nil {
		description = strings.TrimSpace(*req.Description)
	}

	phone := ""
	if req.Phone != nil {
		phone = strings.TrimSpace(*req.Phone)
	}
	address := strings.TrimSpace(req.Address)

	// ==========================================
	// 2. VALIDATE INPUT
	// ==========================================
	if name == "" {
		return nil, apperrors.ErrInvalidRestaurantName
	}

	// ==========================================
	// 3. CREATE RESTAURANT MODEL
	// ==========================================
	restaurant := &models.Restaurant{
		ID:          uuid.New(),
		OwnerID:     ownerID,
		Name:        name,
		Description: &description,
		Phone:       &phone,
		Address:     address,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		IsOpen:      false,
		IsActive:    true,
	}

	// ==========================================
	// 4. BEGIN TRANSACTION
	// ==========================================
	tx, err := s.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, errors.New("failed to begin transaction")
	}

	// ==========================================
	// 5. AUTOMATIC ROLLBACK
	// ==========================================
	defer tx.Rollback()

	// ==========================================
	// 6. CREATE RESTAURANT (use transaction)
	// ==========================================
	err = s.restaurantRepository.CreateTx(
		ctx, tx, restaurant,
	)

	if err != nil {
		return nil, errors.New("failed to create restaurant")
	}

	// ==========================================
	// 7. COMMIT TRANSACTION
	// ==========================================
	if err := tx.Commit(); err != nil {
		return nil, errors.New("failed to commit transaction")
	}

	// ==========================================
	// 8. RETURN RESPONSE
	// ==========================================
	return &dto.RestaurantResponse{
		ID:          restaurant.ID,
		Name:        restaurant.Name,
		Description: restaurant.Description,
		Phone:       restaurant.Phone,
		Address:     restaurant.Address,
		Latitude:    restaurant.Latitude,
		Longitude:   restaurant.Longitude,
		IsOpen:      restaurant.IsOpen,
		IsActive:    restaurant.IsActive,
	}, nil
}

func (s *RestaurantService) GetAllRestaurants(
	ctx context.Context,
) ([]dto.RestaurantResponse, error) {
	// ==========================================
	// 1. GET ALL RESTAURANTS
	// ==========================================
	restaurants, err := s.restaurantRepository.GetAllRestaurants(ctx)

	if err != nil {
		return nil, errors.New("failed to get all restaurant")
	}

	// ==========================================
	// 2. MAP MODEL TO RESPONSE DTO
	// ==========================================
	responses := make([]dto.RestaurantResponse, 0, len(restaurants))

	for _, restaurant := range restaurants {
		response := dto.RestaurantResponse{
			ID:          restaurant.ID,
			Name:        restaurant.Name,
			Description: restaurant.Description,
			Phone:       restaurant.Phone,
			Address:     restaurant.Address,
			Latitude:    restaurant.Latitude,
			Longitude:   restaurant.Longitude,
			IsOpen:      restaurant.IsOpen,
			IsActive:    restaurant.IsActive,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (s *RestaurantService) GetRestaurantByID(
	ctx context.Context,
	id uuid.UUID,
) (*dto.RestaurantResponse, error) {
	// ==========================================
	// 1. GET RESTAURANT BY ID
	// ==========================================
	restaurant, err := s.restaurantRepository.GetRestaurantByID(
		ctx, id,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.ErrRestaurantNotFound
		}

		return nil, errors.New(
			"failed to get restaurant by id",
		)
	}

	// ==========================================
	// 2. MAP MODEL TO RESPONSE DTO
	// ==========================================
	return &dto.RestaurantResponse{
		ID:          restaurant.ID,
		Name:        restaurant.Name,
		Description: restaurant.Description,
		Phone:       restaurant.Phone,
		Address:     restaurant.Address,
		Latitude:    restaurant.Latitude,
		Longitude:   restaurant.Longitude,
		IsOpen:      restaurant.IsOpen,
		IsActive:    restaurant.IsActive,
	}, nil
}

func (s *RestaurantService) SoftDeleteRestaurant(
	ctx context.Context,
	ownerID uuid.UUID,
	restaurantID uuid.UUID,
) error {
	err := s.restaurantRepository.SoftDeleteRestaurant(
		ctx, restaurantID, ownerID,
	)

	if err != nil {
		if errors.Is(err, apperrors.ErrRestaurantNotFound) {
			return apperrors.ErrRestaurantNotFound
		}
		return errors.New("failed to soft delete restaurant")
	}
	return nil
}

func (s *RestaurantService) UpdateRestaurant(
	ctx context.Context,
	ownerID uuid.UUID,
	restaurantID uuid.UUID,
	req dto.UpdateRestaurantRequest,
) error {
	// ==========================================
	// 1. GET RESTAURANT BY ID
	// ==========================================
	restaurant, err := s.restaurantRepository.GetRestaurantByID(
		ctx, restaurantID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return apperrors.ErrRestaurantNotFound
		}
		return errors.New("Failed to get restaurant")
	}

	// ==========================================
	// 2. CHECK OWNER ID
	// ==========================================
	if restaurant.OwnerID != ownerID {
		return apperrors.ErrInvalidRestaurantOwnerID
	}

	// ==========================================
	// 3. VALIDATE AND NORMALIZE INPUT
	// ==========================================
	name := strings.TrimSpace(req.Name)
	phone := strings.TrimSpace(*req.Phone)
	address := strings.TrimSpace(req.Address)

	if name == "" || phone == "" || address == "" {
		return apperrors.ErrInvalidRestaurantData
	}

	// ==========================================
	// 4. UPDATE RESTAURANT
	// ==========================================
	restaurant.Name = name
	restaurant.Description = req.Description
	restaurant.Phone = &phone
	restaurant.Address = address
	restaurant.Latitude = req.Latitude
	restaurant.Longitude = req.Longitude

	err = s.restaurantRepository.UpdateRestaurant(
		ctx, restaurant,
	)
	if err != nil {
		return errors.New("failed to update restaurant")
	}
	return nil
}

func (s *RestaurantService) GetMyRestaurants(
	ctx context.Context,
	ownerID uuid.UUID,
) ([]dto.RestaurantResponse, error) {
	// ==========================================
	// 1. GET RESTAURANTS BY OWNER ID
	// ==========================================
	restaurants, err := s.restaurantRepository.GetMyRestaurants(
		ctx, ownerID,
	)

	if err != nil {
		return nil, errors.New("failed to get all my restaurants")
	}

	// ==========================================
	// 2. MAP MODEL TO RESPONSE DTO
	// ==========================================
	responses := make([]dto.RestaurantResponse, 0, len(restaurants))
	for _, restaurant := range restaurants {
		response := dto.RestaurantResponse{
			ID:          restaurant.ID,
			Name:        restaurant.Name,
			Description: restaurant.Description,
			Phone:       restaurant.Phone,
			Address:     restaurant.Address,
			Latitude:    restaurant.Latitude,
			Longitude:   restaurant.Longitude,
			IsOpen:      restaurant.IsOpen,
			IsActive:    restaurant.IsActive,
		}
		responses = append(responses, response)
	}
	return responses, nil
}

func (s *RestaurantService) UpdateRestaurantStatus(
	ctx context.Context,
	restaurantID uuid.UUID,
	ownerID uuid.UUID,
	isOpen bool,
) error {
	err := s.restaurantRepository.UpdateRestaurantStatus(
		ctx, restaurantID, ownerID, isOpen,
	)
	if err != nil {
		if errors.Is(err, apperrors.ErrRestaurantNotFound) {
			return apperrors.ErrRestaurantNotFound
		}
		return errors.New("failed to update restaurant status")
	}
	return nil
}
