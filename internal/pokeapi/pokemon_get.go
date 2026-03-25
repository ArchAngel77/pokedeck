
package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c Client) GetPokemon(pokemonName string) (Pokemon, error) {
        url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
        if val, ok := c.cache.Get(url); ok {
          pokemonResp := Pokemon{}
          err := json.Unmarshal(val, &pokemonResp)
          if err != nil {
            return Pokemon{}, err
          }
          return pokemonResp, nil
        }
        resp, err := http.Get(url)
        if err != nil {
                return Pokemon{}, err
        }
        defer resp.Body.Close()
        body, err := io.ReadAll(resp.Body)
        if err != nil {
                return Pokemon{}, err
        }
        c.cache.Add(url, body)
        pokemonResp := Pokemon{}
        err = json.Unmarshal(body, &pokemonResp)
        if err != nil {
                return Pokemon{}, err
        }
        return pokemonResp, nil
}
