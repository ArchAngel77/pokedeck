
package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c Client) ListLocations(pageURL *string) (PokemonCity, error) {
	url := "https://pokeapi.co/api/v2/location-area?limit=20"
	if pageURL != nil {
	url = *pageURL
	}
	if val, ok := c.cache.Get(url); ok {
	  pokemonCity := PokemonCity{}
          err := json.Unmarshal(val, &pokemonCity)
          if err != nil {
            return PokemonCity{}, err
          }
          return pokemonCity, nil
	}
	resp, err := http.Get(url)
	if err != nil {
		return PokemonCity{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonCity{}, err
	}
	c.cache.Add(url, body)
	pokemonCity := PokemonCity{}
	err = json.Unmarshal(body, &pokemonCity)
	if err != nil {
		return PokemonCity{}, err
	}
	return pokemonCity, nil
}
