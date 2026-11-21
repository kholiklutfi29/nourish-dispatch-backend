package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/kholiklutfi29/nourish-dispatch/internal/apperrors"
	"github.com/kholiklutfi29/nourish-dispatch/internal/dto"
	"github.com/kholiklutfi29/nourish-dispatch/internal/models"
	"github.com/kholiklutfi29/nourish-dispatch/internal/repository"
)

type MenuService struct {
	db             *sql.DB
	menuRepository *repository.MenuRepository
}

func NewMenuService(
	db *sql.DB,
	menuRepository *repository.MenuRepository,
) *MenuService {
	return &MenuService{
		db:             db,
		menuRepository: menuRepository,
	}
}

func (s *MenuService) CreateMenu(
	ctx context.Context,
	restaurantID uuid.UUID,
	req dto.CreateMenuRequest,
) (*dto.MenuResponse, error) {
	// Normalize input
	name := strings.TrimSpace(req.Name)

	var description *string
	if req.Description != nil {
		value := strings.TrimSpace(*req.Description)
		if value != "" {
			description = &value
		}
	}

	var imageURL *string
	if req.ImageURL != nil {
		value := strings.TrimSpace(*req.ImageURL)
		if value != "" {
			imageURL = &value
		}
	}

	// Validate input
	if name == "" {
		return nil, apperrors.ErrInvalidMenuName
	}

	if req.Price < 0 {
		return nil, apperrors.ErrInvalidMenuPrice
	}

	// create model
	menu := &models.Menu{
		ID:           uuid.New(),
		RestaurantID: restaurantID,
		Name:         name,
		Description:  description,
		Price:        req.Price,
		ImageURL:     imageURL,
		IsAvailable:  true,
	}

	// call repository
	err := s.menuRepository.CreateMenu(
		ctx, menu,
	)

	if err != nil {
		return nil, err
	}

	response := &dto.MenuResponse{
		ID:           menu.ID,
		RestaurantID: menu.RestaurantID,
		Name:         menu.Name,
		Description:  menu.Description,
		Price:        menu.Price,
		ImageURL:     menu.ImageURL,
		IsAvailable:  menu.IsAvailable,
	}

	return response, nil
}

func (s *MenuService) GetMyMenu(
	ctx context.Context,
	restaurantID uuid.UUID,
) ([]dto.MenuResponse, error) {
	// call repository
	menus, err := s.menuRepository.GetMyMenu(
		ctx, restaurantID,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch menu: %w", err)
	}

	// map model to response
	responses := make([]dto.MenuResponse, 0, len(menus))
	for _, menu := range menus {
		response := dto.MenuResponse{
			ID:           menu.ID,
			RestaurantID: menu.RestaurantID,
			Name:         menu.Name,
			Description:  menu.Description,
			Price:        menu.Price,
			ImageURL:     menu.ImageURL,
			IsAvailable:  menu.IsAvailable,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (s *MenuService) GetMenuDetail(
	ctx context.Context,
	menuID uuid.UUID,
) (*dto.MenuResponse, error) {
	// call response
	menu, err := s.menuRepository.GetMenuDetail(
		ctx, menuID,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.ErrMenuNotFound
		}

		return nil, fmt.Errorf("failed to fetch menu detail: %w", err)
	}

	// map model to response
	response := &dto.MenuResponse{
		ID:           menu.ID,
		RestaurantID: menu.RestaurantID,
		Name:         menu.Name,
		Description:  menu.Description,
		Price:        menu.Price,
		ImageURL:     menu.ImageURL,
		IsAvailable:  menu.IsAvailable,
	}

	return response, nil
}

func (s *MenuService) UpdateMenuInfo(
	ctx context.Context,
	menuID uuid.UUID,
	restaurantID uuid.UUID,
	req dto.UpdateMenuInfoRequest,
) error {
	// call repository
	err := s.menuRepository.UpdateMenuInfo(
		ctx, restaurantID, menuID, req.Name, req.Description,
	)

	if err != nil {
		if errors.Is(err, apperrors.ErrMenuNotFound) {
			return apperrors.ErrMenuNotFound
		}
		return fmt.Errorf("failed to update menu info: %w", err)
	}

	return nil

}

func (s *MenuService) UpdateMenuPrice(
	ctx context.Context,
	menuID uuid.UUID,
	restaurantID uuid.UUID,
	req dto.UpdateMenuPriceRequest,
) error {
	// call repository
	err := s.menuRepository.UpdateMenuPrice(
		ctx, restaurantID, menuID, req.Price,
	)

	if err != nil {
		if errors.Is(err, apperrors.ErrMenuNotFound) {
			return apperrors.ErrMenuNotFound
		}
		return fmt.Errorf("failed to update menu price: %w", err)
	}

	return nil
}

func (s *MenuService) UpdateMenuAvailability(
	ctx context.Context,
	menuID uuid.UUID,
	restaurantID uuid.UUID,
	req dto.UpdateMenuAvailabilityRequest,
) error {
	// call repository
	err := s.menuRepository.UpdateMenuAvailability(
		ctx, restaurantID, menuID, req.IsAvailable,
	)

	if err != nil {
		if errors.Is(err, apperrors.ErrMenuNotFound) {
			return apperrors.ErrMenuNotFound
		}
		return fmt.Errorf("failed to update menu availability: %w", err)
	}

	return nil
}

func (s *MenuService) SoftDeleteMenu(
	ctx context.Context,
	menuID uuid.UUID,
	restaurantID uuid.UUID,
) error {
	err := s.menuRepository.SoftDeleteMenu(
		ctx, restaurantID, menuID,
	)

	if err != nil {
		if errors.Is(err, apperrors.ErrMenuNotFound) {
			return apperrors.ErrMenuNotFound
		}
		return fmt.Errorf("failed to update menu availability: %w", err)
	}

	return nil
}

// for customer
func (s *MenuService) GetRestaurantMenus(
	ctx context.Context,
	restaurantID uuid.UUID,
) ([]dto.MenuResponse, error) {
	// call repository
	menus, err := s.menuRepository.GetRestaurantMenus(
		ctx, restaurantID,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch menu: %w", err)
	}

	responses := make([]dto.MenuResponse, 0, len(menus))
	for _, menu := range menus {
		response := dto.MenuResponse{
			ID:           menu.ID,
			RestaurantID: menu.RestaurantID,
			Name:         menu.Name,
			Description:  menu.Description,
			Price:        menu.Price,
			ImageURL:     menu.ImageURL,
			IsAvailable:  menu.IsAvailable,
		}

		responses = append(responses, response)
	}

	return responses, nil
}
