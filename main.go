package main

import (
	"time"

	"github.com/amiulam/pokedex-go/internal/pokeapi"
	"github.com/amiulam/pokedex-go/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokeCache:     *pokeCache,
	}

	startRepl(cfg)
}
