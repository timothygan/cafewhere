package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/timothygan/cafewhere/backend/internal/services"
)

type CafeHandler struct {
	service *services.CoffeeShopService
}

func NewCafeHandler(service *services.CoffeeShopService) *CafeHandler {
	return &CafeHandler{service: service}
}

func (h *CafeHandler) SearchCafes(c *gin.Context) {
	query := c.Query("q")
	lat, _ := strconv.ParseFloat(c.Query("lat"), 64)
	lon, _ := strconv.ParseFloat(c.Query("lon"), 64)

	shops, err := h.service.SearchCafes(c.Request.Context(), query, lat, lon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shops)
}

func (h *CafeHandler) GetCafeDetails(c *gin.Context) {
	id := c.Query("id")

	shop, err := h.service.GetCafeDetails(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shop)
}
