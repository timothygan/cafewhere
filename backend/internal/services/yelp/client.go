package yelp

import (
	"context"

	"github.com/yourusername/coffee-shop-finder-backend/internal/models"
)

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}

func (c *Client) SearchCafes(ctx context.Context, query string, lat, lon float64) ([]*models.Cafe, error) {
	// Implement Yelp API search
	return nil, nil
}
