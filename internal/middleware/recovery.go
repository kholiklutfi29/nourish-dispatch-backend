package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Its primary task is to catch panics (unexpected crashes in Go code) 
// so that your backend server does not go down when a fatal bug occurs.

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func ()  {
			if err := recover(); err != nil {
				log.Printf(
					"panic recovered: %v",
					err,
				)
				
				c.AbortWithStatusJSON(
					500,
					gin.H{
						"success": false,
						"message": "internal server error",
					},
				)

			}
		}()
		c.Next()
	}
}