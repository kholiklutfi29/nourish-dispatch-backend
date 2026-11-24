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

type CartHandler struct {
	cartService *service.CartService
}

func NewCartHandler(
	cartService *service.CartService,
) *CartHandler {
	return &CartHandler{
		cartService: cartService,
	}
}

func (h *CartHandler) CreateCart(c *gin.Context) {

	var req dto.CreateCartRequest

	// get customer id
	customerID, err := helper.JWTIdGetter(c)
	if err != nil {
		return
	}

	// get restaurant id
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(
			c, "Invalid request body",
		)
		return
	}

	// call service
	result, err := h.cartService.CreateCart(
		c.Request.Context(),
		customerID,
		req.RestaurantID,
	)

	if err != nil {
		response.InternalServerError(
			c, "Failed to create cart",
		)
		return
	}

	response.Success(
		c,
		http.StatusCreated,
		"Created cart successfully",
		result,
	)

}

func (h *CartHandler) GetCartByCustomer(c *gin.Context) {
	customerID, err := uuid.Parse(c.Param("customer_id"))
	if err != nil {
		response.BadRequest(c, "Invalid customer ID")
		return
	}

	result, err := h.cartService.GetCartByCustomer(
		c.Request.Context(),
		customerID,
	)
	if err != nil {
		response.InternalServerError(c, "Failed to get cart")
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Customer cart retrieved successfully",
		result,
	)
}

func (h *CartHandler) GetMyCart(c *gin.Context) {
	// get customer id
	customerID, err := helper.JWTIdGetter(c)
	if err != nil {
		return
	}

	// call service
	result, err := h.cartService.GetCartByCustomer(
		c.Request.Context(),
		customerID,
	)

	if err != nil {
		response.InternalServerError(
			c, "Failed to get cart",
		)
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Your cart retrieved successfully",
		result,
	)
}

func (h *CartHandler) DeleteCart(c *gin.Context) {
	// get cart id from param
	cartID, err := uuid.Parse(c.Param("cartID"))
	if err != nil {
		response.BadRequest(c, "Invalid cart ID")
		return
	}

	// get customer id
	customerID, err := helper.JWTIdGetter(c)
	if err != nil {
		return
	}

	// call service
	err = h.cartService.DeleteCart(
		c.Request.Context(),
		customerID,
		cartID,
	)

	if err != nil {
		if errors.Is(err, apperrors.ErrCartNotFound) {
			response.NotFound(c, "Cart not found")
			return
		}

		response.InternalServerError(c, "Failed to get cart")
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Delete Cart Successfully",
		nil,
	)
}
