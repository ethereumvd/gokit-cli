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

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	locationareas := LocationAreaResponse{}
	err = json.Unmarshal(dat, &locationareas)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locationareas, nil
}
