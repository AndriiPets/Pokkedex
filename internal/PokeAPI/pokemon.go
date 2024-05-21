package pokeapi

import (
	"errors"
	"io"
	"log"
	"net/http"
)

var PokemonNotFound = errors.New("Pokemon not found!")

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	if name == "" {
		return Pokemon{}, PokemonNotFound
	}
	url := baseURL + "/pokemon/" + name

	//look in the cache for url
	if val, ok := c.pokeCache.Get(url); ok {
		locs, err := LoadPokemon(val)
		return locs, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	r := resp.StatusCode
	switch {
	case r == 404:
		return Pokemon{}, PokemonNotFound
	case r > 299:
		log.Fatalf("Responce failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	//add to the cache
	c.pokeCache.Add(url, body)

	pokemon, err := LoadPokemon(body)

	return pokemon, err

}

func LoadPokemon(body []byte) (Pokemon, error) {
	locs := &Pokemon{}
	err := ReadJSON(body, locs)

	return *locs, err
}
