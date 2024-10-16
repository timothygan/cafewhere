package api

import (
	"github.com/gorilla/mux"
	"github.com/timothygan/cafewhere/backend/internal/api/handlers"
	"github.com/timothygan/cafewhere/backend/internal/api/middleware"
)

func SetupRoutes(h *handlers.CafeHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/cafes/search", middleware.AuthMiddleware(h.SearchCoffeeShops)).Methods("GET")
	r.HandleFunc("/api/cafes/details", middleware.AuthMiddleware(h.GetCoffeeShopDetails)).Methods("GET")

	return r
}
