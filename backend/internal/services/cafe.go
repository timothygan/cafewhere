package services

import (
	"context"
	"fmt"
	"time"

	"github.com/timothygan/cafewhere/backend/internal/models"
	"github.com/timothygan/cafewhere/backend/internal/repository/cache"
	"github.com/timothygan/cafewhere/backend/internal/repository/postgres"
	"github.com/timothygan/cafewhere/backend/internal/services/osm"
	"github.com/timothygan/cafewhere/backend/internal/utils"
)

type CafeService struct {
	repo        *postgres.CafeRepository
	cacheRepo   *cache.RedisRepository
	osmClient   *osm.Client
	rateLimiter *utils.RateLimiter
}

func NewCafeService(repo *postgres.CafeRepository, cacheRepo *cache.RedisRepository, osmClient *osm.Client, rateLimiter *utils.RateLimiter) *CafeService {
	return &CafeService{
		repo:        repo,
		cacheRepo:   cacheRepo,
		osmClient:   osmClient,
		rateLimiter: rateLimiter,
	}
}

func (s *CafeService) SearchCafes(ctx context.Context, lat, lon float64, radius int) ([]*models.Cafe, error) {
	cacheKey := fmt.Sprintf("search:%f:%f:%d", lat, lon, radius)

	// Try to get from cache
	cachedShops, err := s.cacheRepo.Get(ctx, cacheKey)
	if err == nil {
		return cachedShops, nil
	}

	// If not in cache or error, fetch from OSM
	if s.rateLimiter.Allow() {
		shops, err := s.osmClient.SearchCafes(ctx, lat, lon, radius)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch cafes: %w", err)
		}

		// Cache the results
		if err := s.cacheRepo.Set(ctx, cacheKey, shops, 1*time.Hour); err != nil {
			// Log the error, but don't fail the request
			fmt.Printf("Failed to cache coffee shops: %v", err)
		}

		return shops, nil
	}

	return nil, fmt.Errorf("rate limit exceeded")
}

func (s *CafeService) GetCafeDetails(ctx context.Context, id string) ([]*models.Cafe, error) {
	// First, try to get from cache
	shop, err := s.cacheRepo.Get(ctx, id)
	if err == nil {
		return shop, nil
	}

	// If not in cache, get from database
	shop, err = s.repo.GetCafe(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get coffee shop details: %w", err)
	}

	// Cache the result
	if err := s.cacheRepo.Set(ctx, id, shop, 24*time.Hour); err != nil {
		// Log the error, but don't fail the request
		fmt.Printf("Failed to cache coffee shop details: %v", err)
	}

	return shop, nil
}
