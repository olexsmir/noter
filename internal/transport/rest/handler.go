package rest

import (
	"github.com/Smirnov-O/noter/internal/service"
	v1 "github.com/Smirnov-O/noter/internal/transport/rest/v1"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
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
	handlersV1 := v1.NewHandler(h.services)
	api := r.Group("/api")
	{
		handlersV1.Init(api)
	}
}
