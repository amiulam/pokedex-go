# Pokedex CLI

A command-line interface (CLI) application built in Go that simulates a Pokedex. It interacts with the [PokeAPI](https://pokeapi.co/) to fetch, explore, and catch Pokemon directly from your terminal.

## Features

- **Interactive REPL**: A responsive Read-Eval-Print Loop (REPL) that accepts commands to interact with the Pokemon world.
- **In-Memory Caching**: Implements a concurrent-safe, auto-expiring in-memory cache to minimize API calls and ensure fast responses when revisiting data.
- **Location Exploration**: Paginate through the Pokemon world map and explore specific areas to see what Pokemon live there.
- **Catching Mechanics**: A randomized catch system based on a Pokemon's base experience. Stronger Pokemon are harder to catch!
- **Personal Inventory**: A thread-safe inventory system to store and inspect the Pokemon you've caught.

## Prerequisites

- [Go](https://go.dev/doc/install) 1.20 or higher

## Getting Started

1. Clone the repository (or navigate to the project directory):
   ```bash
   cd pokedex-go
   ```

2. Run the application directly:
   ```bash
   go run .
   ```
   
   *Alternatively, you can build an executable:*
   ```bash
   go build -o pokedex
   ./pokedex
   ```

## Commands

Once the REPL is running, you can use the following commands:

| Command | Description | Example |
| :--- | :--- | :--- |
| `help` | Displays a list of available commands | `help` |
| `map` | Displays the next 20 location areas in the Pokemon world | `map` |
| `mapb` | Displays the previous 20 location areas in the Pokemon world | `mapb` |
| `explore` | Lists all Pokemon located in a specific area. Usage: `explore <area_name>` | `explore canalave-city-area` |
| `catch` | Attempts to catch a specific Pokemon. Usage: `catch <pokemon_name>` | `catch pikachu` |
| `inspect` | View details of a Pokemon you have caught. Usage: `inspect <pokemon_name>` | `inspect pikachu` |
| `pokedex` | Lists all the Pokemon you have caught and stored in your Pokedex | `pokedex` |
| `exit` | Closes the Pokedex and exits the application | `exit` |

## Project Structure

- `main.go` / `repl.go`: Entry point and logic for the REPL environment.
- `command_*.go`: Individual command logic (e.g., map, explore, catch).
- `internal/pokeapi/`: Handles HTTP requests to the PokeAPI and JSON unmarshaling.
- `internal/pokecache/`: Thread-safe caching mechanism with a background garbage collection (reap) loop.
- `internal/pokemon_inventory/`: Thread-safe state management for caught Pokemon.

## License

This project was built as part of the backend development curriculum on [Boot.dev](https://boot.dev).