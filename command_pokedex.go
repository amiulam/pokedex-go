package main

import "fmt"

func commandPokedex(cfg *config) error {
	pokemons := cfg.pokemonInventory.List()

	for _, pokemon := range pokemons {
		fmt.Println("- ", pokemon.Name)
	}

	return nil
}
