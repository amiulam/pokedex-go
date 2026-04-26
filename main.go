package main

import (
	"time"

	"github.com/amiulam/pokedex-go/internal/pokeapi"
	"github.com/amiulam/pokedex-go/internal/pokecache"
	"github.com/amiulam/pokedex-go/internal/pokemon_inventory"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(5 * time.Second)
	pokeInventory := pokemon_inventory.NewPokemonInventory()

	cfg := &config{
		pokeapiClient:    pokeClient,
		pokeCache:        pokeCache,
		pokemonInventory: pokeInventory,
	}

	startRepl(cfg)
}
