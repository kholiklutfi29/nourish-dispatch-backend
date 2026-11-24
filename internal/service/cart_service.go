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

type CartService struct {
	db *sql.DB
	cartRepository *repository.CartRepository
}

func NewCartService(
	db *sql.DB,
	cartRepository *repository.CartRepository,
) *CartService {
	return &CartService{
		db: db,
		cartRepository: cartRepository,
	}
}

func (s *CartService) CreateCart(
	ctx context.Context,
	customerID uuid.UUID,
	restaurantID uuid.UUID,
) (*dto.CartResponse, error) {
	// create model
	cart := &models.Cart{
		ID: uuid.New(),
		CustomerID: customerID,
		RestaurantID: restaurantID,
	}

	// call repository 
	result, err := s.cartRepository.CreateCart(
		ctx, cart,
	)

	if err != nil {
		return nil, err
	}

	response := &dto.CartResponse{
		ID: result.ID,
		CustomerID: result.CustomerID,
		RestaurantID: result.RestaurantID,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	return response, nil
}

func (s *CartService) GetCartByCustomer(
	ctx context.Context,
	costumerID uuid.UUID,
) (*dto.CartResponse, error) {
	// call service
	cart, err := s.cartRepository.GetCartByCustomer(
		ctx, costumerID,
	)

	if err != nil {
		return nil, err
	}

	// parse to dto
	response := &dto.CartResponse{
		ID: cart.ID,
		CustomerID: cart.CustomerID,
		RestaurantID: cart.RestaurantID,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
	}

	return response, nil
}

func (s *CartService) DeleteCart(
	ctx context.Context,
	customerID uuid.UUID,
	cartID uuid.UUID,
) error {
	// call service
	err := s.cartRepository.DeleteCart(
		ctx, customerID, cartID,
	)

	if err != nil {
		if errors.Is(err, apperrors.ErrCartNotFound) {
			return apperrors.ErrCartNotFound
		}
		return fmt.Errorf("failed to update menu availability: %w", err)
	}

	return nil
}