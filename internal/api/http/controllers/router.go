package controllers

import (
	"hey-notes-api/internal/api/http/middleware"
	"hey-notes-api/pkg/services/auth"
	"hey-notes-api/pkg/services/note"

	"github.com/gin-gonic/gin"
)

type RouteImpl struct {
	AuthService auth.AuthService
	NoteService note.NoteService
}

func NewRoute(
	authService auth.AuthService,
	noteService note.NoteService,
) *RouteImpl {
	return &RouteImpl{
		AuthService: authService,
		NoteService: noteService,
	}
}

func (handler *RouteImpl) Route(route *gin.Engine) {
	route.POST("/register", handler.Register)
	route.POST("/login", handler.Login)

	notes := route.Group("/notes")
	notes.Use(middleware.JWTMiddleware())
	
	notes.GET("/", handler.Index)
	notes.GET("/archived", handler.GetArchived)
	notes.POST("/", handler.Create)
	notes.POST("/:slug/archive", handler.Archived)
	notes.POST("/:slug/unarchive", handler.Unarchived)
	notes.DELETE("/:slug", handler.Delete)
}