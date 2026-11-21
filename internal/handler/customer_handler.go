package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kholiklutfi29/nourish-dispatch/internal/apperrors"
	"github.com/kholiklutfi29/nourish-dispatch/internal/dto"
	"github.com/kholiklutfi29/nourish-dispatch/internal/helper"
	"github.com/kholiklutfi29/nourish-dispatch/internal/response"
	"github.com/kholiklutfi29/nourish-dispatch/internal/service"
)

type CustomerHandler struct {
	customerService *service.CustomerService
}

func NewCustomerHandler(
	customerService *service.CustomerService,
) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
	}
}

func (h *CustomerHandler) GetCustomerProfile(c *gin.Context) {
	// get customer id from jwt
	customerID, err := helper.JWTIdGetter(c)

	if err != nil {
		return
	}

	// call service
	customer, err := h.customerService.GetCustomerProfile(
		c.Request.Context(),
		customerID,
	)

	if err != nil {
		if errors.Is(err, apperrors.ErrCustomerNotFound) {
			response.Error(
				c,
				http.StatusNotFound,
				"Customer not found",
			)
			return
		}

		response.Error(
			c,
			http.StatusInternalServerError,
			"failed to fetch customer",
		)
	}

	response.Success(
		c,
		http.StatusOK,
		"Customer fetched successfully",
		customer,
	)
}

func (h *CustomerHandler) GetCostumerAddress(c *gin.Context) {
	// get customer id
	customerID, err := helper.JWTIdGetter(c)

	if err != nil {
		return
	}

	// call service
	customer, err := h.customerService.GetCostumerAddress(
		c.Request.Context(),
		customerID,
	)

	if err != nil {
		if errors.Is(err, apperrors.ErrCustomerNotFound) {
			response.Error(
				c,
				http.StatusNotFound,
				"Customer not found",
			)
			return
		}

		response.Error(
			c,
			http.StatusInternalServerError,
			"failed to fetch customer",
		)
	}

	response.Success(
		c,
		http.StatusOK,
		"Customer address retrieved successfully",
		customer,
	)
}

func (h *CustomerHandler) UpdateCustomerAddress(c *gin.Context) {
	var req dto.CustomerAddressRequest

	// parse json
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(
			c,
			"Invalid request body",
		)
		return
	}

	// get customer id
	customerID, err := helper.JWTIdGetter(c)
	if err != nil {
		return
	}

	// call service
	err = h.customerService.UpdateCustomerAddress(
		c.Request.Context(),
		customerID,
		*req.DefaultAddress,
		*req.Latitude,
		*req.Longitude,
	)

	if err != nil {
		if errors.Is(err, apperrors.ErrCustomerNotFound) {
			response.NotFound(
				c,
				"customer found",
			)
			return
		}
		response.InternalServerError(c, "Failed to update customer")
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Updated customer island successfully",
		nil,
	)

}
