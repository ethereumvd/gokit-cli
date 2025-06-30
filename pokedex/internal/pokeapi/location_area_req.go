package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(nextURL *string) (LocationAreaResponse, error) {

	endpoint := "/location-area"
	fullURL := BaseURL + endpoint
	if nextURL != nil {
		fullURL = *nextURL
	}

	dat, ok := c.cache.Get(fullURL)
	if ok {
		//cache hit
		locationareas := LocationAreaResponse{}
		err := json.Unmarshal(dat, &locationareas)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		return locationareas, nil
	}
	//no cache
	//cache miss

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaResponse{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	locationareas := LocationAreaResponse{}
	err = json.Unmarshal(dat, &locationareas)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	c.cache.AddCache(fullURL, dat)

	return locationareas, nil
}

func (c *Client) ListLocationArea(LocationAreaName string) (LocationArea, error) {

	endpoint := "/location-area/" + LocationAreaName
	fullURL := BaseURL + endpoint

	dat, ok := c.cache.Get(fullURL)
	if ok {
		//cache hit
		locationarea := LocationArea{}
		err := json.Unmarshal(dat, &locationarea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationarea, nil
	}
	//no cache
	//cache miss

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationareas := LocationArea{}
	err = json.Unmarshal(dat, &locationareas)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.AddCache(fullURL, dat)

	return locationareas, nil
}
