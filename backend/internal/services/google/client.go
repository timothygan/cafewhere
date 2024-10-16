package google

import (
	"context"
	"fmt"
	"strings"

	"github.com/timothygan/cafewhere/backend/internal/models"
	"googlemaps.github.io/maps"
)

type PlacesClient struct {
	client *maps.Client
}

func NewPlacesClient(apiKey string) (*PlacesClient, error) {
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create Google Maps client: %w", err)
	}
	return &PlacesClient{client: c}, nil
}

func (pc *PlacesClient) SearchCoffeeShops(ctx context.Context, lat, lng float64, radius uint) ([]*models.Cafe, error) {
	query := fmt.Sprintf("coffee shop")
	r := &maps.FindPlaceFromTextRequest{
		Input:     query,
		InputType: maps.FindPlaceFromTextInputTypeTextQuery,
		LocationBias: maps.LocationBias{
			Circle: &maps.LatLngRadius{
				Location: &maps.LatLng{Lat: lat, Lng: lng},
				Radius:   float64(radius),
			},
		},
		Fields: []maps.PlaceSearchFieldMask{
			maps.PlaceSearchFieldMaskName,
			maps.PlaceSearchFieldMaskFormattedAddress,
			maps.PlaceSearchFieldMaskGeometry,
			maps.PlaceSearchFieldMaskPlaceID,
			maps.PlaceSearchFieldMaskOpeningHours,
			maps.PlaceSearchFieldMaskRating,
			maps.PlaceSearchFieldMaskUserRatingsTotal,
			maps.PlaceSearchFieldMaskPriceLevel,
			maps.PlaceSearchFieldMaskPhotos,
		},
	}

	resp, err := pc.client.FindPlaceFromText(ctx, r)
	if err != nil {
		return nil, fmt.Errorf("failed to perform FindPlaceFromText search: %w", err)
	}

	var shops []*models.Cafe
	for _, candidate := range resp.Candidates {
		shop := &models.Cafe{
			ID:        candidate.PlaceID,
			Name:      candidate.Name,
			Address:   candidate.FormattedAddress,
			Latitude:  candidate.Geometry.Location.Lat,
			Longitude: candidate.Geometry.Location.Lng,
			Rating:    candidate.Rating,
		}
		if candidate.OpeningHours != nil {
			shop.HoursOfOperation = formatOpeningHours(candidate.OpeningHours)
		}
		if len(candidate.Photos) > 0 {
			shop.PhotoURL = getPhotoURL(candidate.Photos[0].PhotoReference, pc.client.APIKey())
		}
		shops = append(shops, shop)
	}

	return shops, nil
}

func (pc *PlacesClient) GetCoffeeShopDetails(ctx context.Context, placeID string) (*models.CoffeeShop, error) {
	r := &maps.PlaceDetailsRequest{
		PlaceID: placeID,
		Fields: []maps.PlaceDetailsFieldMask{
			maps.PlaceDetailsFieldMaskName,
			maps.PlaceDetailsFieldMaskFormattedAddress,
			maps.PlaceDetailsFieldMaskGeometry,
			maps.PlaceDetailsFieldMaskOpeningHours,
			maps.PlaceDetailsFieldMaskRating,
			maps.PlaceDetailsFieldMaskUserRatingsTotal,
			maps.PlaceDetailsFieldMaskPriceLevel,
			maps.PlaceDetailsFieldMaskPhotos,
		},
	}

	resp, err := pc.client.PlaceDetails(ctx, r)
	if err != nil {
		return nil, fmt.Errorf("failed to get place details: %w", err)
	}

	shop := &models.Cafe{
		ID:        resp.PlaceID,
		Name:      resp.Name,
		Address:   resp.FormattedAddress,
		Latitude:  resp.Geometry.Location.Lat,
		Longitude: resp.Geometry.Location.Lng,
		Rating:    resp.Rating,
	}

	if resp.OpeningHours != nil {
		shop.HoursOfOperation = formatOpeningHours(resp.OpeningHours)
	}

	if len(resp.Photos) > 0 {
		shop.PhotoURL = getPhotoURL(resp.Photos[0].PhotoReference, pc.client.APIKey())
	}

	return shop, nil
}

func formatOpeningHours(oh *maps.OpeningHours) string {
	if oh == nil || len(oh.WeekdayText) == 0 {
		return "Hours not available"
	}
	return strings.Join(oh.WeekdayText, ", ")
}

func getPhotoURL(photoReference, apiKey string) string {
	return fmt.Sprintf("https://maps.googleapis.com/maps/api/place/photo?maxwidth=400&photoreference=%s&key=%s", photoReference, apiKey)
}
