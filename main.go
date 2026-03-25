package main

import (
	"github.com/ArchAngel77/pokedeck/internal/pokeapi"
	"time"
)
func main() {
	cfg := &config{
        pokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
	Pokedex:       map[string]pokeapi.Pokemon{},
    }
	startRepl(cfg)
}
