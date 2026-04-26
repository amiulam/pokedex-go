package pokemon_inventory

import (
	"errors"
	"sync"
)

type PokemonInventory struct {
	data map[string]Pokemon
	mu   sync.Mutex
}

type Pokemon struct {
	ID             int
	Name           string
	BaseExperience int
	Height         int
	IsDefault      bool
	Order          int
	Weight         int
}

func NewPokemonInventory() *PokemonInventory {
	return &PokemonInventory{
		data: make(map[string]Pokemon),
	}
}

func (p *PokemonInventory) Add(pokemon Pokemon) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.data[pokemon.Name] = pokemon
}

func (p *PokemonInventory) Get(name string) (Pokemon, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	pokemon, ok := p.data[name]
	if !ok {
		return Pokemon{}, errors.New("No pokemon found")
	}

	return pokemon, nil
}

func (p *PokemonInventory) List() []Pokemon {
	p.mu.Lock()
	defer p.mu.Unlock()
	pokemons := []Pokemon{}
	for _, pokemon := range p.data {
		pokemons = append(pokemons, pokemon)
	}

	return pokemons
}
