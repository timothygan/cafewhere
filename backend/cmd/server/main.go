package main

import (
	"log"
	"net/http"

	"github.com/yourusername/coffee-shop-finder-backend/internal/api"
	"github.com/yourusername/coffee-shop-finder-backend/internal/config"
	"github.com/yourusername/coffee-shop-finder-backend/internal/repository/postgres"
	"github.com/yourusername/coffee-shop-finder-backend/internal/services"
	"github.com/yourusername/coffee-shop-finder-backend/internal/services/osm"
	"github.com/yourusername/coffee-shop-finder-backend/internal/services/yelp"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := postgres.NewConnection(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	repo := postgres.NewCafeRepository(db)
	yelpClient := yelp.NewClient(cfg.YelpAPIKey)
	osmClient := osm.NewClient()

	service := services.NewCoffeeShopService(repo, yelpClient, osmClient)
	handler := api.NewHandler(service)

	router := api.SetupRoutes(handler)

	log.Printf("Starting server on :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
