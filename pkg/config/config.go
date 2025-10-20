package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL      string
	AlpacaAPIKey     string
	AlpacaAPISecret  string
	AlpacaBaseURL    string
}

func Load() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	cfg := &Config{
		DatabaseURL:     os.Getenv("DATABASE_URL"),
		AlpacaAPIKey:    os.Getenv("ALPACA_API_KEY"),
		AlpacaAPISecret: os.Getenv("ALPACA_API_SECRET"),
		AlpacaBaseURL:   os.Getenv("ALPACA_BASE_URL"),
	}

	if cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}

	if cfg.AlpacaAPIKey == "" {
		return nil, fmt.Errorf("ALPACA_API_KEY is required")
	}

	if cfg.AlpacaAPISecret == "" {
		return nil, fmt.Errorf("ALPACA_API_SECRET is required")
	}

	if cfg.AlpacaBaseURL == "" {
		cfg.AlpacaBaseURL = "https://broker-api.sandbox.alpaca.markets"
	}

	return cfg, nil
}
