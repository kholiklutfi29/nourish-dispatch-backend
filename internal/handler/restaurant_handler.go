package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kholiklutfi29/nourish-dispatch/internal/apperrors"
	"github.com/kholiklutfi29/nourish-dispatch/internal/dto"
	"github.com/kholiklutfi29/nourish-dispatch/internal/helper"
	"github.com/kholiklutfi29/nourish-dispatch/internal/response"
	"github.com/kholiklutfi29/nourish-dispatch/internal/service"
)

type RestaurantHandler struct {
	restaurantService *service.RestaurantService
}

func NewRestaurantHandler(restaurantService *service.RestaurantService) *RestaurantHandler {
	return &RestaurantHandler{
		restaurantService: restaurantService,
	}
}

func (h *RestaurantHandler) CreateRestaurant(c *gin.Context) {
	var req dto.CreateRestaurantRequest

	// Parse JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(
			c,
			"Invalid request body",
		)
		return
	}

	// get user id from context
	ownerID, err := helper.JWTIdGetter(c)
	if err != nil {
		return
	}

	// call service
	result, err := h.restaurantService.CreateRestaurant(
		c.Request.Context(),
		ownerID,
		req,
	)
	if err != nil {
		// handle error response
		response.InternalServerError(c, "Failed to create restaurant")
		return
	}

	response.Success(
		c,
		http.StatusCreated,
		"Restaurant Created",
		result,
	)
}

func (h *RestaurantHandler) GetAllRestaurants(c *gin.Context) {
	// call service
	result, err := h.restaurantService.GetAllRestaurants(c.Request.Context())
	if err != nil {
		response.InternalServerError(c, "Failed to get restaurants")
		return
	}

	// return success response
	response.Success(
		c,
		http.StatusOK,
		"Restaurant retrieved successfully",
		result,
	)
}

func (h *RestaurantHandler) GetRestaurantByID(c *gin.Context) {
	// get restaurant id from URL parameter
	restaurantID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid restaurant ID")
	}

	// call service
	result, err := h.restaurantService.GetRestaurantByID(
		c.Request.Context(),
		restaurantID,
	)
	if err != nil {
		if errors.Is(err, apperrors.ErrRestaurantNotFound) {
			response.NotFound(c, "Restaurant not found")
			return
		}

		response.InternalServerError(c, "Failed to get restaurant")
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Restaurant retrieved successfully",
		result,
	)
}

func (h *RestaurantHandler) SoftDeleteRestaurant(c *gin.Context) {
	// get owner id
	ownerID, err := helper.JWTIdGetter(c)
	if err != nil {
		return
	}

	// get restaurant id from URL parameter
	restaurantID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid restaurant ID")
	}

	// call service
	err = h.restaurantService.SoftDeleteRestaurant(
		c.Request.Context(),
		ownerID,
		restaurantID,
	)
	if err != nil {
		if errors.Is(err, apperrors.ErrRestaurantNotFound) {
			response.NotFound(c, "Restaurant not found")
			return
		}

		response.InternalServerError(c, "Failed to get restaurant")
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Restaurant deleted Successfully",
		nil,
	)
}

func (h *RestaurantHandler) UpdateRestaurant(c *gin.Context) {
	var req dto.UpdateRestaurantRequest

	// Parse req
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(
			c,
			"Invalid request body",
		)
		return
	}

	// Get owner id
	ownerID, err := helper.JWTIdGetter(c)
	if err != nil {
		return
	}

	// get restaurant id
	restaurantID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid restaurant ID")
	}

	// call service
	err = h.restaurantService.UpdateRestaurant(
		c.Request.Context(),
		ownerID,
		restaurantID,
		req,
	)

	if err != nil {
		response.InternalServerError(c, "Failed to update restaurant info")
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Update restaurant successfully",
		nil,
	)
}

func (h *RestaurantHandler) GetMyRestaurants(c *gin.Context) {
	// get owner id
	ownerID, err := helper.JWTIdGetter(c)
	if err != nil {
		return
	}

	// call service
	result, err := h.restaurantService.GetMyRestaurants(
		c.Request.Context(),
		ownerID,
	)

	if err != nil {
		response.InternalServerError(c, "Failed to get owner restaurants")
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Restaurant retrieved successfully",
		result,
	)

}

func (h *RestaurantHandler) UpdateRestaurantStatus(c *gin.Context) {
	// parse
	var req dto.UpdateRestaurantStatus

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(
			c,
			"Invalid request body",
		)
		return
	}

	// get owner id
	ownerID, err := helper.JWTIdGetter(c)
	if err != nil {
		return
	}

	// get restaurant id
	restaurantID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return
	}

	// call service
	err = h.restaurantService.UpdateRestaurantStatus(
		c.Request.Context(),
		restaurantID,
		ownerID,
		req.IsOpen,
	)

	if err != nil {
		response.InternalServerError(c, "Failed to update restaurant status")
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Restaurant status update successfully",
		nil,
	)
}
