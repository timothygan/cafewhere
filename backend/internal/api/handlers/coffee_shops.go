package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yourusername/coffee-shop-finder-backend/internal/services"
)

type CafeHandler struct {
	service *services.CoffeeShopService
}

func NewCoffeeShopHandler(service *services.CoffeeShopService) *CafeHandler {
	return &CafeHandler{service: service}
}

func (h *CafeHandler) SearchCoffeeShops(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	lat, _ := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	lon, _ := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)

	shops, err := h.service.SearchCoffeeShops(r.Context(), query, lat, lon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shops)
}

func (h *CafeHandler) GetCoffeeShopDetails(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	shop, err := h.service.GetCoffeeShopDetails(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shop)
}
