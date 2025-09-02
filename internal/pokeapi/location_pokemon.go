package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreaPokemon(name string) (LocationAreaResp, error) {
	url := baseURL + "/location-area/" + name

	// Check cache
	if data, ok := c.cache.Get(url); ok {
		fmt.Println("Cache HIT for", url)
		var resp LocationAreaResp
		if err := json.Unmarshal(data, &resp); err == nil {
			return resp, nil
		}
	} else {
		fmt.Println("Cache MISS for", url)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	// Cache it
	c.cache.Add(url, dat)

	var areaResp LocationAreaResp
	err = json.Unmarshal(dat, &areaResp)
	if err != nil {
		return LocationAreaResp{}, err
	}

	return areaResp, nil
}
