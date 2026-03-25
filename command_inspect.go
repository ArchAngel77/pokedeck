
package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]
	pokemon, ok := cfg.Pokedex[name]
	if !ok {
		return errors.New("You have not caught that pokemon")
	}
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("-%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeinfo := range pokemon.Types {
		fmt.Printf("-%v\n", typeinfo.Type.Name)
	}
	return nil
}
