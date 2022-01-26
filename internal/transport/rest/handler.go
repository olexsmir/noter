package rest

import (
	"fmt"

	"github.com/flof-ik/noter/docs"
	"github.com/flof-ik/noter/internal/config"
	"github.com/flof-ik/noter/internal/service"
	v1 "github.com/flof-ik/noter/internal/transport/rest/v1"
	"github.com/flof-ik/noter/pkg/token"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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
	)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)
	if cfg.Environment != config.EnvLocal {
		docs.SwaggerInfo.Host = cfg.HTTP.Host
	}

	if cfg.Environment != config.Prod {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
