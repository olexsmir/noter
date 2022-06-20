package rest

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/olexsmir/noter/docs"
	"github.com/olexsmir/noter/internal/config"
	"github.com/olexsmir/noter/internal/service"
	v1 "github.com/olexsmir/noter/internal/transport/rest/v1"
	"github.com/olexsmir/noter/pkg/token"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

func (h *Handler) InitRoutes(cfg *config.Config) *gin.Engine {
	r := gin.Default()
	r.Use(
		gin.Recovery(),
		gin.Logger(),
		corsMiddleware,
	)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)
	if cfg.Environment != config.EnvLocal {
		docs.SwaggerInfo.Host = cfg.HTTP.Host
	}

	if cfg.Environment != config.Prod {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

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
