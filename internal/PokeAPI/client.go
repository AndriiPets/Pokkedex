package pokeapi

import (
	"net/http"
	"time"

	"github.com/AndriiPets/pokdex/internal/cache"
)

type Client struct {
	httpClient http.Client
	pokeCache  cache.Cache
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeCache: cache.NewCache(time.Minute * 5),
	}
}
