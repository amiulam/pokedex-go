package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config) error {
	if cfg.pokemonName == "" {
		return errors.New("Please provide the pokemon name. e.g: inspect pikachu")
	}
	
	pokemon, err := cfg.pokemonInventory.Get(cfg.pokemonName)
	if err != nil {
		return err
	}
	
	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	// fmt.Println("Stats: ")
	// fmt.Println("-hp: ", pokemon.Stats[0].Stat.)
	return nil
}
