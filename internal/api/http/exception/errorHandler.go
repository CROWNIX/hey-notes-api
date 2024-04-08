package exception

import (
	"net/http"

	"hey-notes-api/helpers"
	"hey-notes-api/pkg/dto"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Errors *any `json:"errors,omitempty"`
}

type ResponseSuccess struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func HandleError(c *gin.Context, err error) {
	var statusCode int

	switch err.(type) {
	case *NotFound:
		statusCode = http.StatusNotFound
	case *BadRequest:
		statusCode = http.StatusBadRequest
	case *InternalServer:
		statusCode = http.StatusInternalServerError
	case *Unauthorized:
		statusCode = http.StatusUnauthorized
	default:
		statusCode = http.StatusInternalServerError
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: statusCode,
		Message:    err.Error(),
	})

	c.JSON(statusCode, response)
}