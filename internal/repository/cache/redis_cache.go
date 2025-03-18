package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"log/slog"
	"weather-api/internal/core/domain"
)

type RedisCache struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisCache(client *redis.Client, ttl time.Duration) *RedisCache {
	return &RedisCache{client: client, ttl: ttl}
}

func (rc *RedisCache) GetWeather(city string) (*domain.Weather, error) {
	ctx := context.Background()

	val, err := rc.client.Get(ctx, city).Result()
	if err == redis.Nil {
		slog.Debug("Cache miss", "city", city)
		return nil, nil
	} else if err != nil {
		slog.Error("Redis error", "error", err)
		return nil, nil
	}

	var raw map[string]interface{}
	if err := json.Unmarshal([]byte(val), &raw); err != nil {
		slog.Error("Failed to unmarshal cache data", "error", err)
		return nil, err
	}

	// Проверяем наличие и тип каждого поля
	cityName, ok := raw["city"].(string)
	if !ok {
		slog.Warn("Invalid or missing 'city' field in cache data")
		return nil, errors.New("invalid cache data")
	}

	temperatureValue, ok := raw["temperature"].(float64)
	if !ok {
		slog.Warn("Invalid or missing 'temperature' field in cache data")
		return nil, errors.New("invalid cache data")
	}
	temperature := fmt.Sprintf("%.0f°C", temperatureValue)

	description, ok := raw["description"].(string)
	if !ok {
		slog.Warn("Invalid or missing 'description' field in cache data")
		return nil, errors.New("invalid cache data")
	}

	humidity, ok := raw["humidity"].(string)
	if !ok {
		slog.Warn("Invalid or missing 'humidity' field in cache data")
		return nil, errors.New("invalid cache data")
	}

	windSpeed, ok := raw["wind_speed"].(string)
	if !ok {
		slog.Warn("Invalid or missing 'wind_speed' field in cache data")
		return nil, errors.New("invalid cache data")
	}

	return &domain.Weather{
		City:        cityName,
		Temperature: temperature,
		Description: description,
		Humidity:    humidity,
		WindSpeed:   windSpeed,
	}, nil
}

func (rc *RedisCache) SetWeather(city string, weather *domain.Weather) error {
	ctx := context.Background()

	raw := map[string]interface{}{
		"city":        weather.City,
		"temperature": parseTemperatureToFloat(weather.Temperature),
		"description": weather.Description,
		"humidity":    weather.Humidity,
		"wind_speed":  weather.WindSpeed,
	}

	data, err := json.Marshal(raw)
	if err != nil {
		slog.Error("Failed to marshal weather data", "error", err)
		return err
	}

	err = rc.client.Set(ctx, city, data, rc.ttl).Err()
	if err != nil {
		slog.Error("Failed to set cache data", "error", err)
		return err
	}

	slog.Debug("Data cached", "city", city)
	return nil
}

func parseTemperatureToFloat(temp string) float64 {
	temp = strings.ReplaceAll(temp, "°C", "")
	result, _ := strconv.ParseFloat(temp, 64)
	return result
}
