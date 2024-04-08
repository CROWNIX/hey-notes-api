package controllers

import (
	"hey-notes-api/helpers"
	"hey-notes-api/internal/api/http/exception"
	"net/http"

	"hey-notes-api/pkg/dto"

	"github.com/gin-gonic/gin"
)

func (h *RouteImpl) Login(c *gin.Context) {
	payload := new(dto.LoginRequest)

	if err := c.ShouldBindJSON(payload); err != nil {
		exception.HandleError(c,  &exception.BadRequest{Message: err.Error()})
		return
	}
	
	user, err := h.AuthService.Login(c, payload)
	if err != nil {
		exception.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, helpers.Response(dto.ResponseParams{StatusCode: http.StatusOK, Message: "Login Successfully", Data: user}))
}