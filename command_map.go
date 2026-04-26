package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/amiulam/pokedex-go/internal/pokeapi"
)

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.nextLocationURL != nil {
		url = *cfg.nextLocationURL
	}

	var locationResponse pokeapi.LocationAreasResponse
	data, ok := cfg.pokeCache.Get(url)
	if ok {
		err := json.Unmarshal(data, &locationResponse)
		if err != nil {
			return err
		}
	} else {
		resp, err := cfg.pokeapiClient.GetLocationAreas(&url)
		if err != nil {
			return err
		}
		locationResponse = resp

		marshalledData, err := json.Marshal(locationResponse)
		if err != nil {
			return err
		}

		cfg.pokeCache.Add(url, marshalledData)
	}

	cfg.nextLocationURL = locationResponse.Next
	cfg.prevLocationURL = locationResponse.Previous

	for _, location := range locationResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationURL == nil {
		return errors.New("you're on the first page")
	}

	var locationResponse pokeapi.LocationAreasResponse
	data, ok := cfg.pokeCache.Get(*cfg.prevLocationURL)
	if ok {
		err := json.Unmarshal(data, &locationResponse)
		if err != nil {
			return err
		}
	} else {
		resp, err := cfg.pokeapiClient.GetLocationAreas(cfg.prevLocationURL)
		if err != nil {
			return err
		}
		locationResponse = resp

		marshalledData, err := json.Marshal(locationResponse)
		if err != nil {
			return err
		}

		cfg.pokeCache.Add(*cfg.prevLocationURL, marshalledData)
	}

	cfg.nextLocationURL = locationResponse.Next
	cfg.prevLocationURL = locationResponse.Previous

	for _, location := range locationResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}
