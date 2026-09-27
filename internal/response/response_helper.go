package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(
	c *gin.Context,
	status int,
	message string,
	data interface{},
) {
	c.JSON(status, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(
	c *gin.Context,
	status int,
	message string,
) {
	c.JSON(status, Response{
		Success: false,
		Message: message,
	})
}

func BadRequest(
	c *gin.Context,
	message string,
) {
	Error(c, http.StatusBadRequest, message)
}

func Unauthorized(
	c *gin.Context,
	message string,
) {
	Error(c, http.StatusUnauthorized, message)
}

func InternalServerError(
	c *gin.Context,
	message string,
) {
	Error(c, http.StatusInternalServerError, message)
}

// package response

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func Success(c *gin.Context, status int, message string, data interface{}){
// 	c.JSON(status, gin.H{
// 		"success": true,
// 		"message": message,
// 		"data": data,
// 	})
// }

// func Error(c *gin.Context, status int, message string) {
// 	c.JSON(status, gin.H{
// 		"success": false,
// 		"message": message,
// 	})
// }

// func BadRequest(c *gin.Context, message string) {
// 	Error(c, http.StatusBadRequest, message)
// }

// func Unauthorized(c *gin.Context, message string) {
// 	Error(c, http.StatusUnauthorized, message)
// }

// func InternalServerError(c *gin.Context, message string) {
// 	Error(c, http.StatusInternalServerError, message)
// }
