package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocationAreas(pageUrl *string) (LocationAreasResponse, error) {
	url := "https://pokeapi.co/api/v2/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	resp, err := http.Get(url)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationResponse := LocationAreasResponse{}
	err = json.Unmarshal(data, &locationResponse)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	return locationResponse, nil
}
