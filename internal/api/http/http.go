package http

import (
	"fmt"

	"hey-notes-api/internal/api/http/controllers"

	"github.com/gin-gonic/gin"
)

type HttpImpl struct {
	engine    *gin.Engine
	HttpHandler *controllers.RouteImpl
}

func NewHttpImpl(
	httpHandler *controllers.RouteImpl,
) *HttpImpl {
	return &HttpImpl{
		engine: gin.Default(),
		HttpHandler: httpHandler,
	}
}

func (h *HttpImpl) Listen() {
	h.HttpHandler.Route(h.engine)
	
	if err := h.engine.Run(fmt.Sprintf(":%d", 8080)); err != nil {
		panic(err)
	}
}