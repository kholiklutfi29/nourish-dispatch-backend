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

type CartItemService struct {
	db                 *sql.DB
	cartItemRepository *repository.CartItemRepository
	cartRepository *repository.CartRepository
}

func NewCartItemService(
	db *sql.DB,
	cartItemRepository *repository.CartItemRepository,
	cartRepository *repository.CartRepository,
) *CartItemService {
	return &CartItemService{
		db:                 db,
		cartItemRepository: cartItemRepository,
		cartRepository: cartRepository,
	}
}

func (s *CartItemService) CreateCartItem(
	ctx context.Context,
	cartID uuid.UUID,
	menuID uuid.UUID,
	quantity int,
) error {
	// cart item model 
	cartItem := &models.CartItem{
		ID: uuid.New(),
		CartID: cartID,
		MenuID: menuID,
		Quantity: quantity,
	}

	// call repository
	err := s.cartItemRepository.CreateCartItem(
		ctx, cartItem,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *CartItemService) GetCartItems(
	ctx context.Context,
	cartID uuid.UUID,
) ([]dto.CartItemResponse, error) {
	// call repository
	cartItems, err := s.cartItemRepository.GetCartItems(
		ctx, cartID,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch cart items: %w", err)
	}

	// map model to dto
	responses := make([]dto.CartItemResponse, 0, len(cartItems))
	for _, cartItem := range cartItems {
		response := dto.CartItemResponse{
			ID: cartItem.ID,
			CartID: cartItem.CartID,
			MenuID: cartItem.MenuID,
			Quantity: cartItem.Quantity,
			CreatedAt: cartItem.CreatedAt,
			UpdatedAt: cartItem.UpdatedAt,
		}
		responses = append(responses, response)
	}
	return responses, nil
}

func (s *CartItemService) AddCartItem(
    ctx context.Context,
    customerID uuid.UUID,
    request *dto.AddCartItemRequest,
) error {
    cart, err := s.cartRepository.GetCartByCustomer(
        ctx, customerID,
    )
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return apperrors.ErrCartNotFound
        }

        return fmt.Errorf("failed to get customer cart: %w", err)
    }

    cartItem := &models.CartItem{
        ID:       uuid.New(),
        CartID:   cart.ID,
        MenuID:   request.MenuID,
        Quantity: request.Quantity,
    }

    if err := s.cartItemRepository.AddCartItem(ctx, cartItem); err != nil {
        return fmt.Errorf("failed to add cart item: %w", err)
    }

    return nil
}

func (s *CartItemService) DeleteCartItem(
    ctx context.Context,
    customerID uuid.UUID,
    cartItemID uuid.UUID,
) error {
    cart, err := s.cartRepository.GetCartByCustomer(ctx, customerID)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return apperrors.ErrCartNotFound
        }

        return fmt.Errorf("failed to get customer cart: %w", err)
    }

    err = s.cartItemRepository.DeleteCartItem(
        ctx, cartItemID, cart.ID,
    )
    if err != nil {
        if errors.Is(err, apperrors.ErrCartItemNotFound) {
            return apperrors.ErrCartItemNotFound
        }

        return fmt.Errorf("failed to delete cart item: %w", err)
    }

    return nil
}

func (s *CartItemService) ClearCartItems(
    ctx context.Context,
    customerID uuid.UUID,
) error {
    // Get cart
    cart, err := s.cartRepository.GetCartByCustomer(ctx, customerID)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return apperrors.ErrCartNotFound
        }

        return fmt.Errorf("failed to get customer cart: %w", err)
    }

    // Call repository
    if err := s.cartItemRepository.ClearCartItems(
        ctx, cart.ID,
    ); err != nil {
        return fmt.Errorf("failed to clear cart items: %w", err)
    }

    return nil
}

