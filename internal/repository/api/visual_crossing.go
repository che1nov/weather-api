package externalapi

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"weather-api/internal/core/domain"
)

type VisualCrossingAPI struct {
	apiKey string
}

func NewVisualCrossingAPI(apiKey string) *VisualCrossingAPI {
	return &VisualCrossingAPI{apiKey: apiKey}
}

func (api *VisualCrossingAPI) GetWeather(city string) (*domain.Weather, error) {
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=metric&key=%s", city, api.apiKey)
	resp, err := http.Get(url)
	if err != nil {
		slog.Error("Failed to fetch weather data", "error", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Warn("Unexpected status code from API", "status", resp.StatusCode)
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		slog.Error("Failed to decode API response", "error", err)
		return nil, err
	}

	currentConditions := result["currentConditions"].(map[string]interface{})
	temperature := fmt.Sprintf("%.0f°C", currentConditions["temp"].(float64))
	humidity := fmt.Sprintf("%.0f%%", currentConditions["humidity"].(float64))
	windSpeed := fmt.Sprintf("%.0f km/h", currentConditions["windspeed"].(float64))
	description := currentConditions["conditions"].(string)

	return &domain.Weather{
		City:        city,
		Temperature: temperature,
		Description: description,
		Humidity:    humidity,
		WindSpeed:   windSpeed,
	}, nil
}
