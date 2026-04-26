package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type PokemonResponse struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
}

func (c *Client) GetPokemon(name string) (PokemonResponse, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + name

	resp, err := http.Get(url)
	if err != nil {
		return PokemonResponse{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	var pokemonResponse PokemonResponse
	err = json.Unmarshal(data, &pokemonResponse)
	if err != nil {
		return PokemonResponse{}, err
	}

	return pokemonResponse, nil
}
