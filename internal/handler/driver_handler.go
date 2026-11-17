package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kholiklutfi29/nourish-dispatch/internal/apperrors"
	"github.com/kholiklutfi29/nourish-dispatch/internal/dto"
	"github.com/kholiklutfi29/nourish-dispatch/internal/helper"
	"github.com/kholiklutfi29/nourish-dispatch/internal/models"
	"github.com/kholiklutfi29/nourish-dispatch/internal/response"
	"github.com/kholiklutfi29/nourish-dispatch/internal/service"
)

type DriverHandler struct {
	driverService *service.DriverService
}

func NewDriverHandler(
	driverService *service.DriverService,
) *DriverHandler {
	return &DriverHandler{
		driverService: driverService,
	}
}

func (h *DriverHandler) GetDriverProfile(c *gin.Context) {
	// get driver id from jwt middleware
	driverID, err := helper.JWTIdGetter(c)
	if err != nil {
		return
	}

	// call service
	driver, err := h.driverService.GetDriverProfile(
		c.Request.Context(),
		driverID,
	)

	if err != nil {
		if errors.Is(err, apperrors.ErrDriverNotFound) {
			response.Error(
				c,
				http.StatusNotFound,
				"Driver not found",
			)
			return
		}
		response.Error(
			c,
			http.StatusInternalServerError,
			"Failed to fetch driver profile",
		)
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Driver profile fetched successfuly",
		driver,
	)
}

func (h *DriverHandler) UpdateDriverProfile(c *gin.Context) {
	var req dto.UpdateDriverProfileRequest

	// parse json
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(
			c,
			"Invalid request body",
		)
		return
	}

	// get driver id
	driverID, err := helper.JWTIdGetter(c)
	if err != nil {
		return
	}

	// call service
	err = h.driverService.UpdateDriverProfile(
		c.Request.Context(),
		driverID,
		req.VehicleType,
		req.VehiclePlate,
	)

	if err != nil {
		if errors.Is(err, apperrors.ErrDriverNotFound) {
			response.NotFound(
				c,
				"driver not found",
			)
			return
		}
		response.InternalServerError(c, "Failed to update driver profile")
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Driver profile updated successfuly",
		nil,
	)
}

func (h *DriverHandler) UpdateDriverStatus(c *gin.Context) {
	var req dto.UpdateDriverStatusRequest

	// bind
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(
			c,
			"Invalid request body",
		)
		return
	}

	// get driver id
	driverID, err := helper.JWTIdGetter(c)
	if err != nil {
		return
	}

	// call service
	err = h.driverService.UpdateDriverStatus(
		c.Request.Context(),
		driverID,
		models.DriverStatus(req.Status),
	)

	if err != nil {
		if errors.Is(err, apperrors.ErrInvalidDriverStatus) {
			response.NotFound(
				c,
				"Invalid Driver Status",
			)
			return
		}
		response.InternalServerError(c, "Failed to update driver status")
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Update driver status successfully",
		nil,
	)
}

func (h *DriverHandler) GetAllDrivers(c *gin.Context) {
	// call service
	result, err := h.driverService.GetAllDrivers(
		c.Request.Context(),
	)

	if err != nil {
		response.InternalServerError(
			c,
			"Failed to fetch drivers",
		)
		return
	}

	// return fetch data
	response.Success(
		c, 
		http.StatusOK,
		"Fetching drivers successfully",
		result,
	)
}

func (h *DriverHandler) GetDriverByStatus(c *gin.Context) {
	var req dto.UpdateDriverStatusRequest

	// bind
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(
			c,
			"Invalid request body",
		)
		return
	}

	// call service
	result, err := h.driverService.GetDriverByStatus(
		c.Request.Context(),
		models.DriverStatus(req.Status),
	)

	if err != nil {
		log.Printf("GetDriverByStatus error: %v", err)

		response.InternalServerError(
			c,
			"Failed to fetch drivers",
		)
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Driver retrieved successfully",
		result,
	)
}

func (h *DriverHandler) GetDriverDetails(c *gin.Context) {

	// get driver id from jwt middleware
	driverID, err := helper.JWTIdGetter(c)
	if err != nil {
		return
	}

	// call service
	driver, err := h.driverService.GetDriverDetails(
		c.Request.Context(),
		driverID,
	)

	if err != nil {
		response.InternalServerError(
			c,
			"Failed to fetch drivers",
		)
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Fetched driver detail successfully",
		driver,
	)
}