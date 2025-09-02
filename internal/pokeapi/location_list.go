package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil && *pageURL != "" {
		url = *pageURL
	}

	if data, ok := c.cache.Get(url); ok {
        fmt.Println("Cache HIT for", url)
        var resp RespShallowLocations
        if err := json.Unmarshal(data, &resp); err == nil {
            return resp, nil
        }
        // If unmarshal fails, treat as cache miss and fall through
    } else {
        fmt.Println("Cache MISS for", url)
    }

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}



	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, dat)

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}


	return locationsResp, nil
}