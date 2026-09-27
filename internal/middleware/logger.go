package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func (c *gin.Context) {

		start := time.Now()

		//  this line temporarily halts code execution within 
		// the Logger() function and hands control over to the
		//  main Handler/Controller (such as the Register or Login function) 
		// to process the request to completion.
		c.Next()
		// If you do not call `c.Next()`: the middleware will calculate 
		// the duration and `c.Writer.Status()` before the controller
		// has a chance to process anything.


		duration := time.Since(start)

		log.Printf(
			"%s %s %d %v",
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			duration,
		)
	}
}