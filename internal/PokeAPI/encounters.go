package pokeapi

import (
	"errors"
	"io"
	"net/http"
)

var AreaNotFound = errors.New("Area not found!")

func (c *Client) GetEncounters(area string) (ExploreAreaTable, error) {
	url := baseURL + "/location-area/" + area

	//look in the cache for url
	if val, ok := c.pokeCache.Get(url); ok {
		locs, err := LoadEncounters(val)
		return locs, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ExploreAreaTable{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ExploreAreaTable{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return ExploreAreaTable{}, AreaNotFound
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ExploreAreaTable{}, err
	}

	//add to the cache
	c.pokeCache.Add(url, body)

	encs, err := LoadEncounters(body)

	return encs, err
}

func LoadEncounters(body []byte) (ExploreAreaTable, error) {
	locs := &ExploreAreaTable{}
	err := ReadJSON(body, locs)

	return *locs, err
}
