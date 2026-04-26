package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/amiulam/pokedex-go/internal/pokeapi"
	"github.com/amiulam/pokedex-go/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient   pokeapi.Client
	pokeCache       pokecache.Cache
	nextLocationURL *string
	prevLocationURL *string
	locationName    string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		text := scanner.Text()
		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		if len(words) > 1 {
			cfg.locationName = words[1]
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
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display next page of location area",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous page of location area",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Display detail of location area: explore <area_name>",
			callback:    commandExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
