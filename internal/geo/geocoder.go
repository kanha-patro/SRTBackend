package geo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Geocoder struct {
	APIKey string
}

type GeocodeResponse struct {
	Results []GeocodeResult `json:"results"`
}

type GeocodeResult struct {
	FormattedAddress string `json:"formatted_address"`
	Geometry         struct {
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
	} `json:"geometry"`
}

func NewGeocoder(apiKey string) *Geocoder {
	return &Geocoder{APIKey: apiKey}
}

func (g *Geocoder) Geocode(address string) (float64, float64, error) {
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", address, g.APIKey)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, 0, fmt.Errorf("failed to geocode address: %s", resp.Status)
	}

	var geocodeResponse GeocodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&geocodeResponse); err != nil {
		return 0, 0, err
	}

	if len(geocodeResponse.Results) == 0 {
		return 0, 0, fmt.Errorf("no results found for address: %s", address)
	}

	location := geocodeResponse.Results[0].Geometry.Location
	return location.Lat, location.Lng, nil
}
