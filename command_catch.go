package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"

	"github.com/amiulam/pokedex-go/internal/pokeapi"
	"github.com/amiulam/pokedex-go/internal/pokemon_inventory"
)

func commandCatch(cfg *config) error {
	if cfg.pokemonName == "" {
		return errors.New("Please provide the pokemon name. e.g: catch pikachu")
	}

	msg := fmt.Sprintf("Throwing a Pokeball at %s...", cfg.pokemonName)
	fmt.Println(msg)
	fmt.Println()

	url := "https://pokeapi.co/api/v2/pokemon/" + cfg.pokemonName

	var pokemonResp pokeapi.PokemonResponse
	data, ok := cfg.pokeCache.Get(url)
	if ok {
		err := json.Unmarshal(data, &pokemonResp)
		if err != nil {
			return err
		}
	} else {
		resp, err := cfg.pokeapiClient.GetPokemon(cfg.pokemonName)
		if err != nil {
			return err
		}
		pokemonResp = resp

		marshalledData, err := json.Marshal(pokemonResp)
		if err != nil {
			return err
		}

		cfg.pokeCache.Add(url, marshalledData)
	}

	const threshold = 40
	randNum := rand.Intn(pokemonResp.BaseExperience)

	if randNum > threshold {
		fmt.Println(pokemonResp.Name, "escaped!")
		return nil
	}

	fmt.Println(pokemonResp.Name, "was caught!")

	cfg.pokemonInventory.Add(pokemon_inventory.Pokemon{
		ID:             pokemonResp.ID,
		Name:           pokemonResp.Name,
		BaseExperience: pokemonResp.BaseExperience,
		Height:         pokemonResp.Height,
		Weight:         pokemonResp.Weight,
	})

	return nil
}
