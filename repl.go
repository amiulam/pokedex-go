package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/amiulam/pokedex-go/internal/pokeapi"
	"github.com/amiulam/pokedex-go/internal/pokecache"
	"github.com/amiulam/pokedex-go/internal/pokemon_inventory"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	pokeCache        *pokecache.Cache
	pokemonInventory *pokemon_inventory.PokemonInventory
	nextLocationURL  *string
	prevLocationURL  *string
	locationName     string
	pokemonName      string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		if len(words) > 1 {
			cfg.locationName = words[1]
			cfg.pokemonName = words[1]
		} else {
			cfg.locationName = ""
			cfg.pokemonName = ""
		}

		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	return strings.Fields(lower)
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a list of available commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Lists all Pokemon located in a specific area. Usage: explore <area_name>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a specific Pokemon. Usage: catch <pokemon_name>",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "View details of a Pokemon you have caught. Usage: inspect <pokemon_name>",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all the Pokemon you have caught and stored in your Pokedex",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Closes the Pokedex and exits the application",
			callback:    commandExit,
		},
	}
}
