package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kholiklutfi29/nourish-dispatch/internal/response"
)

func JWTIdGetter(c *gin.Context) (uuid.UUID, error) {
	// get user id from context
	userID, exist := c.Get("user_id")
	if !exist {
		response.Unauthorized(
			c,
			"User id not found",
		)
		return uuid.Nil, errors.New("User ID not found")
	}

	userIDString, ok := userID.(string)
	if !ok {
		response.Unauthorized(c, "Invalid user ID")
		return uuid.Nil, errors.New("Invalid user ID type")
	}

	userUUID, err := uuid.Parse(userIDString)
	if err != nil {
		response.Unauthorized(c, "Invalid user id")
		return uuid.Nil, errors.New("Invalid user ID format")
	}

	return userUUID, nil
}