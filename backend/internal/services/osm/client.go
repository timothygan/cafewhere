package osm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/timothygan/cafewhere/backend/internal/models"
	"github.com/timothygan/cafewhere/backend/internal/utils"
	"net/http"
	"strings"
)

const (
	overpassURL = "https://overpass-api.de/api/interpreter"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

func (c *Client) SearchCafes(ctx context.Context, lat, lon float64, radius int) ([]*models.Cafe, error) {
	query := fmt.Sprintf(`data=[out:json];(node["amenity"="cafe"](around:%d,%f,%f);way["amenity"="cafe"](around:%d,%f,%f);relation["amenity"="cafe"](around:%d,%f,%f););out center;`,
		radius, lat, lon, radius, lat, lon, radius, lat, lon,
	)

	req, err := http.NewRequestWithContext(ctx, "POST", overpassURL, strings.NewReader(query))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	fmt.Println(req.URL.String())
	fmt.Println(req.Header)
	fmt.Println(req.Body)
	fmt.Println(query)
	fmt.Println()
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Overpass API returned non-OK status: %d", resp.StatusCode)
	}

	var result struct {
		Elements []struct {
			Type string  `json:"type"`
			ID   int64   `json:"id"`
			Lat  float64 `json:"lat"`
			Lon  float64 `json:"lon"`
			Tags struct {
				Name               string `json:"name"`
				Amenity            string `json:"amenity"`
				Opening            string `json:"opening_hours"`
				LastCheckedOpening string `json:"check_date:opening_hours"`
				Wifi               string `json:"internet_access"`
				Operator           string `json:"operator"`
			} `json:"tags"`
		} `json:"elements"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	shops := make([]*models.Cafe, 0, len(result.Elements))
	for _, element := range result.Elements {
		if element.Tags.Amenity == "cafe" {
			var openHours models.WeekOpeningHours
			if _openHours, err := utils.ParseOpeningHours(element.Tags.Opening); err != nil {
				fmt.Printf("failed to parse opening hours: %v", err)
			} else {
				openHours = *_openHours
			}
			shop := &models.Cafe{
				ID:               fmt.Sprintf("osm:%d", element.ID),
				Name:             element.Tags.Name,
				Latitude:         element.Lat,
				Longitude:        element.Lon,
				HoursOfOperation: openHours,
				HoursLastUpdated: element.Tags.LastCheckedOpening,
				HasWifi:          element.Tags.Wifi == "yes" || element.Tags.Wifi == "wlan",
				IsIndependent:    element.Tags.Operator == "",
			}
			if shop.IsIndependent {
				fmt.Println(shop.Name)
				fmt.Println(shop.HasWifi)
				fmt.Println(shop.HoursLastUpdated)
				fmt.Println(shop.HoursOfOperation)
			}
			shops = append(shops, shop)
		}
	}

	return shops, nil
}
