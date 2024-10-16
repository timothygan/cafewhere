package services

import (
	"context"
	"fmt"
	"time"

	"github.com/timothygan/cafewhere/backend/internal/models"
	"github.com/timothygan/cafewhere/backend/internal/repository/cache"
	"github.com/timothygan/cafewhere/backend/internal/repository/postgres"
	"github.com/timothygan/cafewhere/backend/internal/services/osm"
	"github.com/timothygan/cafewhere/backend/internal/services/yelp"
	"github.com/timothygan/cafewhere/backend/internal/utils"
)

type CoffeeShopService struct {
	repo        *postgres.CafeRepository
	cacheRepo   *cache.RedisRepository
	yelpClient  *yelp.Client
	osmClient   *osm.Client
	rateLimiter *utils.RateLimiter
}

func NewCoffeeShopService(repo *postgres.CafeRepository, cacheRepo *cache.RedisRepository, yelpClient *yelp.Client, osmClient *osm.Client, rateLimiter *utils.RateLimiter) *CoffeeShopService {
	return &CoffeeShopService{
		repo:        repo,
		cacheRepo:   cacheRepo,
		yelpClient:  yelpClient,
		osmClient:   osmClient,
		rateLimiter: rateLimiter,
	}
}

func (s *CoffeeShopService) SearchCafes(ctx context.Context, query string, lat, lon float64) ([]*models.Cafe, error) {
	cacheKey := fmt.Sprintf("search:%s:%f:%f", query, lat, lon)
	cachedShops, err := s.cacheRepo.Get(ctx, cacheKey)
	if err == nil {
		return cachedShops, nil
	}

	var shops []*models.Cafe
	if s.rateLimiter.Allow() {
		shops, err = s.yelpClient.SearchCafes(ctx, query, lat, lon)
		if err != nil {
			// Log the error
			shops, err = s.osmClient.SearchCafes(ctx, query, lat, lon)
			if err != nil {
				return nil, err
			}
		}
	} else {
		shops, err = s.osmClient.SearchCafes(ctx, query, lat, lon)
		if err != nil {
			return nil, err
		}
	}

	s.cacheRepo.Set(ctx, cacheKey, shops, 1*time.Hour)
	return shops, nil
}

func (s *CoffeeShopService) GetCafeDetails(ctx context.Context, id string) (*models.Cafe, error) {
	cacheKey := fmt.Sprintf("shop:%s", id)
	cachedShop, err := s.cacheRepo.Get(ctx, cacheKey)
	if err == nil {
		return cachedShop, nil
	}

	shop, err := s.repo.GetCafe(ctx, id)
	if err != nil {
		return nil, err
	}

	s.cacheRepo.Set(ctx, cacheKey, shop, 24*time.Hour)
	return shop, nil
}
