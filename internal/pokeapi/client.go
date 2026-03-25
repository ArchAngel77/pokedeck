package pokeapi

import (
	"net/http"
	"time"
	"github.com/ArchAngel77/pokedeck/internal/pokecache"
)

type Client struct {
    httpClient http.Client
    cache      *pokecache.Cache
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
