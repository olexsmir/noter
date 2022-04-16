package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/olexsmir/noter/internal/service"
	"github.com/olexsmir/noter/pkg/token"
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
