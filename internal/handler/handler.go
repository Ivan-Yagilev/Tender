package handler

import (
	"github.com/gin-gonic/gin"
	"tender/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	filter := router.Group("/")
	{
		filter.GET("/kpgz/:kpgz", h.getAllKpgz)
		filter.GET("/inn/:inn", h.getProviderByInn)
	}

	return router
}
