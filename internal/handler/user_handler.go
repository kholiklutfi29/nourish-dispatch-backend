package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kholiklutfi29/nourish-dispatch/internal/apperrors"
	"github.com/kholiklutfi29/nourish-dispatch/internal/dto"
	"github.com/kholiklutfi29/nourish-dispatch/internal/response"
	"github.com/kholiklutfi29/nourish-dispatch/internal/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) ChangeUserPhone(c *gin.Context) {
	// ==========================================
	// 1. GET USER ID FROM CONTEXT
	// ==========================================
	userID, exist := c.Get("user_id") // get function return any, need to parse

	if !exist {
		response.Unauthorized(
			c,
			"User id not found",
		)
		return
	}

	// type assertion from any to string
	userIDString, ok := userID.(string)

	if !ok {
		response.Unauthorized(c, "invalid user id")
		return
	}

	// ==========================================
	// 2. PARSE REQUEST BODY
	// ==========================================
	var req dto.UserChangePhoneRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(
			c,
			"Invalid request body",
		)
		return
	}

	// ==========================================
	// 3. CALL SERVICE
	// ==========================================
	result, err := h.userService.ChangeUserPhone(c, req, userIDString)

	if err != nil {
		log.Printf("change user phone error: %v", err)
		response.InternalServerError(c, "failed to update phone")
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Update phone succeed",
		result,
	)

}

func (h *UserHandler) ChangeUserPassword(c *gin.Context) {
	// ==========================================
	// 1. GET USER ID FROM CONTEXT
	// ==========================================

	userID, exist := c.Get("user_id")

	if !exist {
		response.Unauthorized(
			c,
			"User id not found",
		)
		return
	}

	// type assertion from any to string
	userIDString, ok := userID.(string)

	if !ok {
		response.Unauthorized(c, "invalid user id")
		return
	}

	// ==========================================
	// 2. PARSE REQUEST BODY
	// ==========================================
	var req dto.UserChangePasswordRequest 

	if err:= c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(
			c,
			"Invalid request body",
		)
		return
	}

	// ==========================================
	// 3. CALL SERVICE
	// ==========================================
	err := h.userService.ChangeUserPassword(
		c,
		req,
		userIDString,
	)

	if err != nil {
		if errors.Is(err, apperrors.ErrInternal) {
			response.InternalServerError(
				c,
				"Internal server error",
			)
			return
		}

		response.BadRequest(c, err.Error())
		return
	}	

	response.Success(
		c,
		http.StatusOK,
		"Update Password Succeed",
		nil,
	)
}

func (h *UserHandler) ChangeUserName(c *gin.Context) {

	// ==========================================
	// 1. GET USER ID FROM CONTEXT
	// ==========================================
	userID, exist := c.Get("user_id") // get function return any, need to parse

	if !exist {
		response.Unauthorized(
			c,
			"User id not found",
		)
		return
	}

	// type assertion from any to string
	userIDString, ok := userID.(string)

	if !ok {
		response.Unauthorized(c, "invalid user id")
		return
	}

	// ==========================================
	// 2. PARSE REQUEST BODY
	// ==========================================
	var req dto.UserChangeNameRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(
			c,
			"Invalid request body",
		)
		return
	}

	// ==========================================
	// 3. CALL SERVICE
	// ==========================================
	result, err := h.userService.ChangeUserName(c, req, userIDString)

	if err != nil {
		log.Printf("change user name error: %v", err)
		response.InternalServerError(c, "failed to update name")
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Update name succeed",
		result,
	)
}
