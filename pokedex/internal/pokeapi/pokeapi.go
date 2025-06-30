package pokeapi

import (
	"net/http"
	"time"

	"github.com/ethereumvd/gokit-cli/pokedex/internal/pokecache"
)

const BaseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(CacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(CacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
