package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"log/slog"
	"weather-api/internal/config"
	"weather-api/internal/core/usecase"
	h "weather-api/internal/delivery/http"
	externalapi "weather-api/internal/repository/api"
	"weather-api/internal/repository/cache"
)

func main() {
	config.SetupLogger()

	cfg, err := config.LoadConfig(".env")
	if err != nil {
		slog.Error("Failed to load config", "error", err)
		return
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})

	ctx := context.Background()
	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		slog.Error("Failed to connect to Redis", "error", err)
		return
	}
	slog.Info("Connected to Redis", "response", pong)

	weatherRepo := externalapi.NewVisualCrossingAPI(cfg.VisualCrossingAPIKey)
	cacheRepo := cache.NewRedisCache(redisClient, 12*time.Hour)

	weatherUsecase := usecase.NewWeatherUseCase(weatherRepo, cacheRepo)

	weatherHandler := h.NewWeatherHandler(weatherUsecase)

	http.HandleFunc("/weather", weatherHandler.GetWeather)

	slog.Info("Starting server", "address", ":8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
