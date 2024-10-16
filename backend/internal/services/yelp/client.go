package yelp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/timothygan/cafewhere/backend/internal/models"
)

const (
	baseURL      = "https://api.yelp.com/v3"
	searchPath   = "/businesses/search"
	businessPath = "/businesses/%s"
)

type Client struct {
	apiKey     string
	httpClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}
}

func (c *Client) SearchCafes(ctx context.Context, query string, lat, lon float64) ([]*models.Cafe, error) {
	endpoint, err := url.Parse(baseURL + searchPath)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	params := url.Values{}
	params.Add("term", "cafe")
	params.Add("latitude", strconv.FormatFloat(lat, 'f', 6, 64))
	params.Add("longitude", strconv.FormatFloat(lon, 'f', 6, 64))
	params.Add("radius", "1000") // 1km radius
	params.Add("limit", "20")    // Fetch 20 results
	endpoint.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", endpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Yelp API returned non-OK status: %d", resp.StatusCode)
	}

	var searchResult struct {
		Businesses []struct {
			ID          string  `json:"id"`
			Name        string  `json:"name"`
			Rating      float64 `json:"rating"`
			ReviewCount int     `json:"review_count"`
			Coordinates struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"coordinates"`
			Location struct {
				Address1 string `json:"address1"`
				City     string `json:"city"`
				Country  string `json:"country"`
			} `json:"location"`
			Price      string `json:"price"`
			ImageURL   string `json:"image_url"`
			Categories []struct {
				Alias string `json:"alias"`
			} `json:"categories"`
		} `json:"businesses"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&searchResult); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	cafes := make([]*models.Cafe, 0, len(searchResult.Businesses))
	for _, business := range searchResult.Businesses {
		// Check if it's actually a cafe
		isCafe := false
		for _, category := range business.Categories {
			if category.Alias == "coffee" || category.Alias == "cafes" {
				isCafe = true
				break
			}
		}
		if !isCafe {
			continue
		}

		shop := &models.Cafe{
			ID:        business.ID,
			Name:      business.Name,
			Address:   fmt.Sprintf("%s, %s, %s", business.Location.Address1, business.Location.City, business.Location.Country),
			Latitude:  business.Coordinates.Latitude,
			Longitude: business.Coordinates.Longitude,
			Rating:    business.Rating,
			PhotoURL:  business.ImageURL,
		}
		cafes = append(cafes, shop)
	}

	return cafes, nil
}

func (c *Client) GetCafeDetails(ctx context.Context, id string) (*models.Cafe, error) {
	endpoint, err := url.Parse(fmt.Sprintf(baseURL+businessPath, id))
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", endpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Yelp API returned non-OK status: %d", resp.StatusCode)
	}

	var businessDetails struct {
		ID          string  `json:"id"`
		Name        string  `json:"name"`
		Rating      float64 `json:"rating"`
		ReviewCount int     `json:"review_count"`
		Coordinates struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"coordinates"`
		Location struct {
			Address1 string `json:"address1"`
			City     string `json:"city"`
			Country  string `json:"country"`
		} `json:"location"`
		Price    string `json:"price"`
		ImageURL string `json:"image_url"`
		Hours    []struct {
			Open []struct {
				Start string `json:"start"`
				End   string `json:"end"`
				Day   int    `json:"day"`
			} `json:"open"`
		} `json:"hours"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&businessDetails); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	shop := &models.Cafe{
		ID:        businessDetails.ID,
		Name:      businessDetails.Name,
		Address:   fmt.Sprintf("%s, %s, %s", businessDetails.Location.Address1, businessDetails.Location.City, businessDetails.Location.Country),
		Latitude:  businessDetails.Coordinates.Latitude,
		Longitude: businessDetails.Coordinates.Longitude,
		Rating:    businessDetails.Rating,
		PhotoURL:  businessDetails.ImageURL,
	}

	// Process hours of operation
	if len(businessDetails.Hours) > 0 {
		shop.HoursOfOperation = formatHours(businessDetails.Hours[0].Open)
	}

	return shop, nil
}

func formatHours(openHours []struct {
	Start string `json:"start"`
	End   string `json:"end"`
	Day   int    `json:"day"`
}) string {
	// Implement logic to format hours of operation
	// This is a placeholder and should be implemented based on your needs
	return "Hours formatting not implemented"
}
