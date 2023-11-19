package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	filter := router.Group("/filter")
	{
		filter.GET("/kpgz/:kpgz", h.getAllKpgz)
		filter.GET("/inn/:inn", h.getProviderByInn)
	}

	return router
}
