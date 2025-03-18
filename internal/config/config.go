package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisAddr            string
	VisualCrossingAPIKey string
}

func LoadConfig(envPath string) (*Config, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		log.Printf("Warning: Failed to load .env file: %v", err)
	}

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		return nil, fmt.Errorf("missing environment variable: REDIS_ADDR")
	}

	apiKey := os.Getenv("VISUAL_CROSSING_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("missing environment variable: VISUAL_CROSSING_API_KEY")
	}

	return &Config{
		RedisAddr:            redisAddr,
		VisualCrossingAPIKey: apiKey,
	}, nil
}
