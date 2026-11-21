package dto

import (
	"github.com/google/uuid"
)

type CreateMenuRequest struct {
	Name         string    `json:"name"`
	Description  *string   `json:"description"`
	Price        float64   `json:"price"`
	ImageURL     *string   `json:"image_url"`
}

type MenuResponse struct {
	ID           uuid.UUID `json:"id"`
	RestaurantID uuid.UUID `json:"restaurant_id"`
	Name         string    `json:"name"`
	Description  *string   `json:"description"`
	Price        float64   `json:"price"`
	ImageURL     *string   `json:"image_url"`
	IsAvailable  bool      `json:"is_available"`
}

type UpdateMenuInfoRequest struct {
	Name         string    `json:"name"`
	Description  string   `json:"description"`
}

type UpdateMenuPriceRequest struct {
	Price        float64   `json:"price"`
}

type UpdateMenuAvailabilityRequest struct {
	IsAvailable  bool      `json:"is_available"`
}


