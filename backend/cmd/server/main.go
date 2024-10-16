package main

import (
	"github.com/timothygan/cafewhere/backend/internal/api/handlers"
	"log"
	"net/http"

	"github.com/timothygan/cafewhere/backend/internal/api"
	"github.com/timothygan/cafewhere/backend/internal/config"
	"github.com/timothygan/cafewhere/backend/internal/repository/postgres"
	"github.com/timothygan/cafewhere/backend/internal/services"
	"github.com/timothygan/cafewhere/backend/internal/services/osm"
	"github.com/timothygan/cafewhere/backend/internal/services/yelp"
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

	service := services.NewCafeService(repo, yelpClient, osmClient)
	handler := handlers.NewHandler(service)

	router := api.SetupRoutes(handler)

	log.Printf("Starting server on :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
