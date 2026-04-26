package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/amiulam/pokedex-go/internal/pokemon_inventory"
)

func commandCatch(cfg *config) error {
	if cfg.pokemonName == "" {
		return errors.New("Please provide the pokemon name. e.g: catch pikachu")
	}

	msg := fmt.Sprintf("Throwing a Pokeball at %s...", cfg.pokemonName)
	fmt.Println(msg)
	fmt.Println()

	resp, err := cfg.pokeapiClient.GetPokemon(cfg.pokemonName)
	if err != nil {
		return err
	}

	const threshold = 40
	randNum := rand.Intn(resp.BaseExperience)

	if randNum > threshold {
		fmt.Println(resp.Name, "escaped!")
		return nil
	}

	fmt.Println(resp.Name, "was caught!")

	cfg.pokemonInventory.Add(pokemon_inventory.Pokemon{
		ID:             resp.ID,
		Name:           resp.Name,
		BaseExperience: resp.BaseExperience,
		Height:         resp.Height,
		Weight:         resp.Weight,
	})

	return nil
}
