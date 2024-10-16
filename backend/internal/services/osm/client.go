package osm

import (
	"context"

	"github.com/yourusername/coffee-shop-finder-backend/internal/models"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) SearchCafes(ctx context.Context, query string, lat, lon float64) ([]*models.Cafe, error) {
	// Implement OpenStreetMap search
	return nil, nil
}
