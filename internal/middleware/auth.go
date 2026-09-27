package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kholiklutfi29/nourish-dispatch/internal/jwt"
	"github.com/kholiklutfi29/nourish-dispatch/internal/response"
)

func Auth(jwtService *jwt.JWTService) gin.HandlerFunc {

	return func(c *gin.Context) {

		// get header
		// example: Authorization: Bearer eyJhbGciOiJIUzI1NiIs...
		//  authHeader contains: Bearer eyJhbGciOiJIUzI1NiIs...
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" { // case if header nil
			response.Unauthorized(
				c,
				"Authorization header required",
			)

			c.Abort() // stop chain middleware/handler after this middleware
			return // out from this function
		}

		// split based on space blank :" ", and maks split was 2 parts
		// example : parts[0] = "Bearer", parts[1] = "abcdef123456"
		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			response.Unauthorized(
				c,
				"invalid authorization header",
			)

			c.Abort()
			return
		}

		tokenString := parts[1]

		claims, err := jwtService.ValidateJWTToken(tokenString)

		if err != nil {
			response.Unauthorized(
				c,
				"invalid or expired token",
			)

			c.Abort()
			return
		}

		// Save information user to Gin Context
		c.Set("user_id", claims.UserID)
		c.Set("user_role", claims.Role)

		c.Next()
	}

}