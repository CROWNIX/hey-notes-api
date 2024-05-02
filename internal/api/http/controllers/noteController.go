package controllers

import (
	"hey-notes-api/helpers"
	"hey-notes-api/internal/api/http/exception"
	"net/http"

	"hey-notes-api/pkg/dto"

	"github.com/gin-gonic/gin"
)

func (h *RouteImpl) Index(c *gin.Context) {
	notes, err := h.NoteService.GetAllNotes(c)
	if err != nil {
		exception.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, helpers.Response(dto.ResponseParams{StatusCode: http.StatusOK, Message: "Success", Data: notes}))
}

func (h *RouteImpl) GetArchived(c *gin.Context) {
	notes, err := h.NoteService.GetArchivedNotes(c)
	if err != nil {
		exception.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, helpers.Response(dto.ResponseParams{StatusCode: http.StatusOK, Message: "Archived", Data: notes}))
}

func (h *RouteImpl) Create(c *gin.Context) {
	payload := new(dto.NoteRequest)

	if err := c.ShouldBindJSON(payload); err != nil {
		exception.HandleError(c,  &exception.BadRequest{Message: err.Error()})
		return
	}
	
	note, err := h.NoteService.Create(c, payload)
	if err != nil {
		exception.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, helpers.Response(dto.ResponseParams{StatusCode: http.StatusCreated, Message: "Note Has Been Created", Data: note}))
}

func (h *RouteImpl) Archived(c *gin.Context) {
	slug := c.Param("slug")
	
	note, err := h.NoteService.Archived(c, slug)
	if err != nil {
		exception.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, helpers.Response(dto.ResponseParams{StatusCode: http.StatusOK, Message: "Note Has Been Archived", Data: note}))
}

func (h *RouteImpl) Show(c *gin.Context) {
	slug := c.Param("slug")
	
	note, err := h.NoteService.FindBySlug(c, slug)
	if err != nil {
		exception.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, helpers.Response(dto.ResponseParams{StatusCode: http.StatusOK, Message: "Success", Data: note}))
}

func (h *RouteImpl) Unarchived(c *gin.Context) {
	slug := c.Param("slug")
	
	note, err := h.NoteService.Unarchived(c, slug)
	if err != nil {
		exception.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, helpers.Response(dto.ResponseParams{StatusCode: http.StatusOK, Message: "Note Has Been Unarchived", Data: note}))
}

func (h *RouteImpl) Delete(c *gin.Context) {
	slug := c.Param("slug")
	
	note, err := h.NoteService.Delete(c, slug)
	if err != nil {
		exception.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, helpers.Response(dto.ResponseParams{StatusCode: http.StatusOK, Message: "Note Has Been Deleted", Data: note}))
}