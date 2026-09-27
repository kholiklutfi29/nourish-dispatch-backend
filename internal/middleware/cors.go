package middleware

import "github.com/gin-gonic/gin"

// Granting access permissions and configuring security rules to allow 
// frontend applications (such as Web or Flutter) from other domains or 
// ports to communicate with the backend API.

func CORS() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set(
			"Access-Control-Allow-Origin",
			"*", // set to your domain (* means allow all)
		)

		c.Writer.Header().Set(
			"Access-Control-Allow-Methods",
			"GET, POST, PUT, DELETE, PATCH, OPTIONS",
		)

		c.Writer.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Authorization",
		)

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()

	}

}