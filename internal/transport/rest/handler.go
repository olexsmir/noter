package rest

import (
	"github.com/flof-ik/noter/internal/service"
	v1 "github.com/flof-ik/noter/internal/transport/rest/v1"
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

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	h.initApi(r)

	return r
}

func (h *Handler) initApi(r *gin.Engine) {
	handlersV1 := v1.NewHandler(h.services, h.tokenManager)
	api := r.Group("/api")
	{
		handlersV1.Init(api)
	}
}
