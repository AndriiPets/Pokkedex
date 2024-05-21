package pokeapi

import (
	"io"
	"log"
	"net/http"
)

func (c *Client) GetLocations(pageURL *string) (LocationsTable, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	//look in the cache for url
	if val, ok := c.pokeCache.Get(url); ok {
		locs, err := LoadLocations(val)
		return locs, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationsTable{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationsTable{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		log.Fatalf("Responce failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationsTable{}, err
	}

	//add to the cache
	c.pokeCache.Add(url, body)

	locs, err := LoadLocations(body)
	return locs, err
}

func LoadLocations(body []byte) (LocationsTable, error) {
	locs := &LocationsTable{}
	err := ReadJSON(body, locs)

	return *locs, err
}
