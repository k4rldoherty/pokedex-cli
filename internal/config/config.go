package config

import (
	"time"

	"github.com/k4rldoherty/pokedex-cli/internal/api"
)

type Config struct {
	Client   api.Client
	Next     *string
	Previous *string
}

func NewConfig(cacheInterval time.Duration) *Config {
	return &Config{
		Client: api.NewClient(cacheInterval),
	}
}
