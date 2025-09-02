package pokeapi

import (
	"net/http"
	"time"

	"github.com/shahrishi501/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
			
		},
		cache: pokecache.NewCache(5 * time.Minute),
	}
}
