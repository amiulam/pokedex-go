package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	locationResponse, err := cfg.pokeapiClient.GetLocationArea(cfg.nextLocationURL)
	if err != nil {
		return err
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

	locationResponse, err := cfg.pokeapiClient.GetLocationArea(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationResponse.Next
	cfg.prevLocationURL = locationResponse.Previous

	for _, location := range locationResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}
