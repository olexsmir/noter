package v1

import (
	"github.com/flof-ik/noter/internal/service"
	"github.com/flof-ik/noter/pkg/token"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services     *service.Services
	tokenManager token.TokenManager
}

func NewHandler(services *service.Services, tokenManager token.TokenManager) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initAuthRoutes(v1)
		h.initNotebooksRoutes(v1)
	}
}
