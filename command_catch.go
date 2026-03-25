

package main

import (
        "errors"
	"fmt"
        "math/rand"
)

func commandCatch(cfg *config, args ...string) error {
        if len(args) == 0 {
                return errors.New("you must provide a Pokemon name")
        }
        name := args[0]
	fmt.Printf("Throwing a Pokeball at %v...\n", name)
        pokemon, err := cfg.pokeapiClient.GetPokemon(name)
        if err != nil {
                return err
	}
	randomNumber := rand.Intn(600)
	if randomNumber < pokemon.BaseExperience {
			fmt.Printf("%v escaped!\n", name)
	} else {
		fmt.Printf("%v was caught!\n", name)
		cfg.Pokedex[name] = pokemon
        }
        return nil
}

