package controllers

import (
	"hey-notes-api/helpers"
	"hey-notes-api/internal/api/http/exception"
	"net/http"

	"hey-notes-api/pkg/dto"

	"github.com/gin-gonic/gin"
)

func (h *RouteImpl) Register(c *gin.Context) {
	payload := new(dto.RegisterRequest)

	if err := c.ShouldBindJSON(payload); err != nil {
		exception.HandleError(c,  &exception.BadRequest{Message: err.Error()})
		return
	}
	
	user, err := h.AuthService.Register(c, payload)
	if err != nil {
		exception.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, helpers.Response(dto.ResponseParams{StatusCode: http.StatusCreated, Message: "Register Successfully", Data: user}))
}