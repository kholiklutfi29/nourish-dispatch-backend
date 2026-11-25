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

type CartItemHandler struct {
	cartItemService *service.CartItemService
	cartService     *service.CartService
}

func NewCartItemHandler(
	cartItemService *service.CartItemService,
	cartService *service.CartService,
) *CartItemHandler {
	return &CartItemHandler{
		cartItemService: cartItemService,
		cartService:     cartService,
	}
}

func (h *CartItemHandler) CreateCartItem(c *gin.Context) {
	var req dto.CreateCartItemRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(
			c, "Invalid request body",
		)
		return
	}

	// get customer id
	customerID, err := helper.JWTIdGetter(c)
	if err != nil {
		return
	}

	// get cart id by customer id
	cart, err := h.cartService.GetCartByCustomer(
		c.Request.Context(),
		customerID,
	)

	if err != nil {
		if errors.Is(err, apperrors.ErrCartNotFound) {
			response.NotFound(c, "Cart not found")
			return
		}

		response.InternalServerError(c, "Failed to get customer cart")
		return
	}

	// parse menu id
	menuID, err := uuid.Parse(req.MenuID)
	if err != nil {
		response.BadRequest(c, "Invalid menu ID")
		return
	}

	// call service
	err = h.cartItemService.CreateCartItem(
		c.Request.Context(),
		cart.ID,
		menuID,
		req.Quantity,
	)

	if err != nil {
		response.InternalServerError(
			c, "Failed to create cart item",
		)
		return
	}

	response.Success(
		c,
		http.StatusCreated,
		"Successfully created cart item",
		nil,
	)
}

func (h *CartItemHandler) GetCartItems(c *gin.Context) {
	// get customer id
	customerID, err := helper.JWTIdGetter(c)
	if err != nil {
		return
	}

	// get cart id by customer id
	cart, err := h.cartService.GetCartByCustomer(
		c.Request.Context(),
		customerID,
	)

	if err != nil {
		if errors.Is(err, apperrors.ErrCartNotFound) {
			response.NotFound(c, "Cart not found")
			return
		}

		response.InternalServerError(c, "Failed to get customer cart")
		return
	}

	// call service
	result, err := h.cartItemService.GetCartItems(
		c.Request.Context(), cart.ID,
	)

	if err != nil {
		response.InternalServerError(
			c, "Failed to fetch cart items",
		)
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Cart items fetched successfully",
		result,
	)
}

func (h *CartItemHandler) AddCartItem(c *gin.Context) {
    var req dto.AddCartItemRequest

    // Get customer ID
    customerID, err := helper.JWTIdGetter(c)
    if err != nil {
        return
    }

    // Bind request body
    if err := c.ShouldBindJSON(&req); err != nil {
        response.BadRequest(c, "Invalid request body")
        return
    }

    // Call service
    err = h.cartItemService.AddCartItem(
        c.Request.Context(),
        customerID,
        &req,
    )
    if err != nil {
        if errors.Is(err, apperrors.ErrCartNotFound) {
            response.NotFound(c, "Cart not found")
            return
        }

        response.InternalServerError(c, "Failed to add cart item")
        return
    }

    response.Success(
        c,
        http.StatusOK,
        "Cart item added successfully",
        nil,
    )
}

func (h *CartItemHandler) DeleteCartItem(c *gin.Context) {
	// Get customer ID
    customerID, err := helper.JWTIdGetter(c)
    if err != nil {
        return
    }

	// Parse cart item ID
    cartItemID, err := uuid.Parse(c.Param("cartItemID"))
    if err != nil {
        response.BadRequest(c, "Invalid cart item ID")
        return
    }

	// call service
	err = h.cartItemService.DeleteCartItem(
		c.Request.Context(), customerID, cartItemID,
	)

	if err != nil {
        if errors.Is(err, apperrors.ErrCartNotFound) {
            response.NotFound(c, "Cart not found")
            return
        }

        if errors.Is(err, apperrors.ErrCartItemNotFound) {
            response.NotFound(c, "Cart item not found")
            return
        }

        response.InternalServerError(c, "Failed to delete cart item")
        return
    }

	response.Success(
		c,
		http.StatusOK,
		"Cart item deleted successfully",
		nil,
	)

}

func (h *CartItemHandler) ClearCartItems(c *gin.Context) {
    // Get customer ID
    customerID, err := helper.JWTIdGetter(c)
    if err != nil {
        return
    }

    // Call service
    err = h.cartItemService.ClearCartItems(
        c.Request.Context(),
        customerID,
    )
    if err != nil {
        if errors.Is(err, apperrors.ErrCartNotFound) {
            response.NotFound(c, "Cart not found")
            return
        }

        response.InternalServerError(c, "Failed to clear cart items")
        return
    }

    response.Success(
        c,
        http.StatusOK,
        "Cart items cleared successfully",
        nil,
    )
}