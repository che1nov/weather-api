package usecase

import (
	"errors"
	"log"
	"weather-api/internal/core/domain"
)

type WeatherCache interface {
	GetWeather(city string) (*domain.Weather, error)
	SetWeather(city string, weather *domain.Weather) error
}

type WeatherFetcher interface {
	GetWeather(city string) (*domain.Weather, error)
}

type WeatherUseCase struct {
	apiFetcher WeatherFetcher
	cache      WeatherCache
}

func NewWeatherUseCase(apiFetcher WeatherFetcher, cache WeatherCache) *WeatherUseCase {
	return &WeatherUseCase{
		apiFetcher: apiFetcher,
		cache:      cache,
	}
}

func (uc *WeatherUseCase) GetWeather(city string) (*domain.Weather, error) {
	if city == "" {
		return nil, errors.New("city cannot be empty")
	}

	weather, err := uc.cache.GetWeather(city)
	if err != nil {
		return nil, err
	}
	if weather != nil {
		return weather, nil
	}

	weather, err = uc.apiFetcher.GetWeather(city)
	if err != nil {
		return nil, err
	}

	if err := uc.cache.SetWeather(city, weather); err != nil {
		log.Printf("Failed to cache weather data: %v", err)
	}

	return weather, nil
}
