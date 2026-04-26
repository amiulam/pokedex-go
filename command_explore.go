package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/amiulam/pokedex-go/internal/pokeapi"
)

func commandExplore(cfg *config) error {
	if cfg.locationName == "" {
		return errors.New("please provide the location name. e.g: explore canalave-city-area")
	}

	url := "https://pokeapi.co/api/v2/location-area/" + cfg.locationName

	msg := fmt.Sprintf("Exploring %s...", cfg.locationName)
	fmt.Println()
	fmt.Println(msg)

	var locationResponse pokeapi.LocationAreaResponse
	data, ok := cfg.pokeCache.Get(url)

	if ok {
		err := json.Unmarshal(data, &locationResponse)
		if err != nil {
			return err
		}
	} else {
		resp, err := cfg.pokeapiClient.GetLocationArea(cfg.locationName)
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

	for _, item := range locationResponse.PokemonEncounters {
		fmt.Println("-", item.Pokemon.Name)
	}

	cfg.locationName = ""

	return nil
}
