package config

import (
	"time"

	"github.com/k4rldoherty/pokedex-cli/internal/api"
	"github.com/k4rldoherty/pokedex-cli/internal/pokedex"
)

type Config struct {
	Client   api.Client
	PokeDex  pokedex.Pokedex
	Next     *string
	Previous *string
}

func NewConfig(cacheInterval time.Duration) *Config {
	return &Config{
		Client:  api.NewClient(cacheInterval),
		PokeDex: *pokedex.NewPokedex(),
	}
}
