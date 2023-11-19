package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tender/internal/entity"
)

type getAllKpgzResponse struct {
	Data []entity.ProviderResponse `json:"data"`
}

func (h *Handler) getAllKpgz(c *gin.Context) {
	kpgz := c.Param("kpgz")

	lists, err := h.services.GetAllKpgz(kpgz)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllKpgzResponse{
		Data: lists,
	})
}

func (h *Handler) getProviderByInn(c *gin.Context) {
	inn, err := strconv.Atoi(c.Param("inn"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid inn param")
		return
	}

	providerStat, err := h.services.GetProviderByInn(inn)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, providerStat)
}
