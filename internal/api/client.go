package api

import (
	"net/http"
	"time"

	"github.com/k4rldoherty/pokedex-cli/internal/cache"
)

type Client struct {
	cache      cache.Cache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache:      *cache.NewCache(cacheInterval),
		httpClient: http.Client{},
	}
}
