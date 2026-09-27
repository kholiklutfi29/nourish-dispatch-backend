package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kholiklutfi29/nourish-dispatch/internal/apperrors"
	"github.com/kholiklutfi29/nourish-dispatch/internal/dto"
	"github.com/kholiklutfi29/nourish-dispatch/internal/response"
	"github.com/kholiklutfi29/nourish-dispatch/internal/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {

	var req dto.RegisterRequest

	// Parse JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(
			c,
			"Invalid request body",
		)
		return
	}

	// Service
	result, err := h.authService.Register(
		c.Request.Context(),
		req,
	)

	if err != nil {

		if errors.Is(
			err,
			apperrors.ErrEmailAlreadyRegistered,
		) {
			response.Error(
				c,
				http.StatusConflict,
				err.Error(),
			)
			return
		}

		response.InternalServerError(
			c,
			err.Error(),
		)
		return
	}

	response.Success(
		c,
		http.StatusCreated,
		"Registration successful",
		result,
	)
}

func (h *AuthHandler) Login(c *gin.Context) {

	var req dto.LoginRequest

	// Parse JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(
			c,
			"Invalid request body",
		)
		return
	}

	// Service
	result, err := h.authService.Login(
		c.Request.Context(),
		req,
	)

	if err != nil {
		response.Unauthorized(
			c,
			err.Error(),
		)
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Login successful",
		result,
	)
}

// func (h *AuthHandler) Register(c *gin.Context) {

// 	// c *gin.Context was receive data from http post frontend
	
// 	var req dto.RegisterRequest

// 	// parse json (ShouldBindJSON read stream data http body json that is stored in c *gin.Context)
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		response.BadRequest(
// 			c,
// 			"Invalid request body",
// 		)
// 		return
// 	}

// 	// service
// 	result, err := h.authService.Register(
// 		c.Request.Context(),
// 		req,
// 	)

// 	if err != nil {
// 		if errors.Is(err, apperrors.ErrEmailAlreadyRegistered) {
// 			response.Error(
// 				c,
// 				http.StatusConflict,
// 				err.Error(),
// 			)
// 		}

// 		// handling internal server error
// 		response.InternalServerError(
// 			c,
// 			err.Error(),
// 		)
// 	}

// 	response.Success(
// 		c,
// 		http.StatusCreated,
// 		"Registration successful",
// 		result,
// 	)
// }

// func (h *AuthHandler) Login(c *gin.Context) {

// 	var req dto.LoginRequest

// 	// parse json
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		response.BadRequest(
// 			c,
// 			"Invalid request body",
// 		)
// 		return
// 	}

// 	// Service
// 	result, err := h.authService.Login(
// 		c.Request.Context(),
// 		req,
// 	)

// 	if err != nil {
// 		response.Unauthorized(
// 			c,
// 			err.Error(),
// 		)
// 	}

// 	response.Success(
// 		c,
// 		http.StatusOK,
// 		"Login successful",
// 		result,
// 	)

// }