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

type MenuHandler struct {
	menuService       *service.MenuService
	restaurantService *service.RestaurantService
}

func NewMenuHandler(
	menuService *service.MenuService,
	restaurantService *service.RestaurantService,
) *MenuHandler {
	return &MenuHandler{
		menuService:       menuService,
		restaurantService: restaurantService,
	}
}

func (h *MenuHandler) CreateMenu(c *gin.Context) {

	var req dto.CreateMenuRequest

	// parse
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(
			c, "Invalid request body",
		)
		return
	}

	// get authenticated user id
	ownerID, err := helper.JWTIdGetter(c)
	if err != nil {
		return
	}

	// get restaurant id owned by merchant
	restaurantID, err := h.restaurantService.GetRestaurantIDByOwner(
		c.Request.Context(),
		ownerID,
	)

	if err != nil {
		if errors.Is(err, apperrors.ErrRestaurantNotFound) {
			response.NotFound(c, "Restaurant not found")
			return
		}

		response.InternalServerError(c, "Failed to retrieve restaurant")
		return
	}

	// call service
	result, err := h.menuService.CreateMenu(
		c.Request.Context(), restaurantID, req,
	)

	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrInvalidMenuName),
			errors.Is(err, apperrors.ErrInvalidMenuPrice):
			response.BadRequest(
				c, err.Error(),
			)

		default:
			response.InternalServerError(
				c, "Failed to create menu",
			)
		}
		return
	}

	response.Success(
		c,
		http.StatusCreated,
		"Created menu successfully",
		result,
	)
}

func (h *MenuHandler) GetMyMenu(c *gin.Context) {
	// get restaurant id from jwt
	ownerID, err := helper.JWTIdGetter(c)
	if err != nil {
		return
	}

	// call service
	result, err := h.menuService.GetMyMenu(
		c.Request.Context(), ownerID,
	)

	if err != nil {
		response.InternalServerError(c, "Failed to get your restaurant menu")
		return 
	}

	response.Success(
		c,
		http.StatusOK,
		"Your restaurant menu retrieved successfully",
		result,
	)
}

func (h *MenuHandler) GetMenuDetail(c *gin.Context) {
	// get menu id
	menuID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid menu ID")
		return
	}

	// call service
	result, err := h.menuService.GetMenuDetail(
		c.Request.Context(), menuID,
	)

	if err != nil {
		if errors.Is(err, apperrors.ErrMenuNotFound) {
			response.NotFound(c, "Menu not found")
			return
		}

		response.InternalServerError(c, "Failed to get menu detail")
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Menu detail retrieved successfully",
		result,
	)
}

func (h *MenuHandler) UpdateMenuInfo(c *gin.Context) {
	var req dto.UpdateMenuInfoRequest

	// parse
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(
			c,
			"Invalid request body",
		)
	}

	// get restaurantID
	ownerID, err := helper.JWTIdGetter(c)
	if err != nil {
		return
	}

	// get menuID from param
	menuID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid menu ID")
		return
	}

	// call service
	err = h.menuService.UpdateMenuInfo(
		c.Request.Context(),
		menuID,
		ownerID,
		req,
	)


	if err != nil {
		if errors.Is(err, apperrors.ErrMenuNotFound) {
			response.NotFound(c, "Menu not found")
			return
		}

		response.InternalServerError(c, "Failed to update menu info")
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Update menu info successfuly",
		nil,
	)

}

func (h *MenuHandler) UpdateMenuPrice(c *gin.Context) {
	var req dto.UpdateMenuPriceRequest

	// parse
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(
			c,
			"Invalid request body",
		)
	}

	// get id from params
	menuID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid menu ID")
		return
	}

	// get restaurantID
	ownerID, err := helper.JWTIdGetter(c)
	if err != nil {
		return
	}

	// call service
	err = h.menuService.UpdateMenuPrice(
		c.Request.Context(),
		menuID,
		ownerID,
		req,
	)

	if err != nil {
		if errors.Is(err, apperrors.ErrMenuNotFound) {
			response.NotFound(c, "Menu not found")
			return
		}

		response.InternalServerError(c, "Failed to update menu price")
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Successfully update menu price",
		nil,
	)
}

func (h *MenuHandler) UpdateMenuAvailability(c *gin.Context) {
	// should bind
	var req dto.UpdateMenuAvailabilityRequest

	// parse
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(
			c,
			"Invalid request body",
		)
	}

	// get id from params
	menuID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid menu ID")
		return
	}

	// get restaurantID
	ownerID, err := helper.JWTIdGetter(c)
	if err != nil {
		return
	}

	// call service
	err = h.menuService.UpdateMenuAvailability(
		c.Request.Context(),
		menuID,
		ownerID,
		req,
	)

	if err != nil {
		if errors.Is(err, apperrors.ErrMenuNotFound) {
			response.NotFound(c, "Menu not found")
			return
		}

		response.InternalServerError(c, "Failed to update menu availability")
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Successfully update menu availability",
		nil,
	)
}

func (h *MenuHandler) SoftDeleteMenu(c *gin.Context) {
	// get id from params
	menuID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid menu ID")
		return
	}

	// get restaurantID
	ownerID, err := helper.JWTIdGetter(c)
	if err != nil {
		return
	}

	// call service
	err = h.menuService.SoftDeleteMenu(
		c.Request.Context(),
		menuID,
		ownerID,
	)

	if err != nil {
		if errors.Is(err, apperrors.ErrMenuNotFound) {
			response.NotFound(c, "Menu not found")
			return
		}

		response.InternalServerError(c, "Failed to delete menu")
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Successfully deleted menu",
		nil,
	)
}

func (h *MenuHandler) GetRestaurantMenus(c *gin.Context) {
	restaurantID, err := uuid.Parse(c.Param("restaurantID"))
	if err != nil {
		response.BadRequest(c, "Invalid restaurant ID")
		return
	}

	// call service
	result, err := h.menuService.GetRestaurantMenus(
		c.Request.Context(),
		restaurantID,
	)

	if err != nil {
		response.InternalServerError(c, "Failed to get restaurant menu")
		return 
	}

	response.Success(
		c,
		http.StatusOK,
		"Successfully fetch restaurant menus",
		result,
	)
}