package main

import (
	"github.com/timothygan/cafewhere/backend/internal/api/handlers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/timothygan/cafewhere/backend/internal/api"
	"github.com/timothygan/cafewhere/backend/internal/config"
	"github.com/timothygan/cafewhere/backend/internal/repository/cache"
	"github.com/timothygan/cafewhere/backend/internal/repository/postgres"
	"github.com/timothygan/cafewhere/backend/internal/services"
	"github.com/timothygan/cafewhere/backend/internal/services/osm"
	"github.com/timothygan/cafewhere/backend/internal/services/yelp"
	"github.com/timothygan/cafewhere/backend/internal/utils"
	"go.uber.org/zap"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Sync()

	// Set up database connection
	db, err := postgres.NewConnection(cfg.DatabaseURL)
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}
	defer db.Close()

	// Set up Redis connection
	redisClient := redis.NewClient(&redis.Options{
		Addr: cfg.RedisURL,
	})
	defer redisClient.Close()

	// Initialize repositories
	repo := postgres.NewCafeRepository(db)
	cacheRepo := cache.NewRedisRepository(redisClient)

	// Initialize clients and services
	yelpClient := yelp.NewClient(cfg.YelpAPIKey)
	osmClient := osm.NewClient()
	rateLimiter := utils.NewRateLimiter(5, 1) // 5 requests per second

	service := services.NewCoffeeShopService(repo, cacheRepo, yelpClient, osmClient, rateLimiter)
	handler := handlers.NewCoffeeShopHandler(service)

	// Set up Gin
	r := gin.New()
	r.Use(gin.Recovery())

	// Set up routes
	api.SetupRoutes(r, handler, logger)

	// Start server
	logger.Info("Starting server", zap.String("port", cfg.Port))
	if err := r.Run(":" + cfg.Port); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}
