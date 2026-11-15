package dto

import "github.com/google/uuid"

type CreateRestaurantRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description *string  `json:"description"`
	Phone       *string  `json:"phone"`
	Address     string   `json:"address" binding:"required"`
	Latitude    *float64 `json:"latitude"`
	Longitude   *float64 `json:"longitude"`
}

type RestaurantResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Phone       *string   `json:"phone"`
	Address     string    `json:"address"`
	Latitude    *float64  `json:"latitude"`
	Longitude   *float64  `json:"longitude"`
	IsOpen      bool      `json:"is_open"`
	IsActive    bool      `json:"is_active"`
}

type UpdateRestaurantStatusRequest struct {
	IsOpen bool `json:"is_open"`
}

type UpdateRestaurantRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description *string  `json:"description"`
	Phone       *string  `json:"phone"`
	Address     string   `json:"address" binding:"required"`
	Latitude    *float64 `json:"latitude"`
	Longitude   *float64 `json:"longitude"`
}

type UpdateRestaurantStatus struct {
	IsOpen bool `json:"is_open"`
}
