
package main

import (
	"fmt"
	"errors"
)

func commandMap(cfg *config, args ...string) error {
	locations, err := cfg.pokeapiClient.ListLocations(cfg.next)
	if err != nil {
		return err
	}
	cfg.next = locations.Next
	cfg.previous = locations.Previous
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapBack(cfg *config, args ...string) error {
	 if cfg.previous == nil {
                return errors.New("you're on the first page")
        }
	locations, err := cfg.pokeapiClient.ListLocations(cfg.previous)
        if err != nil {
                return err
        }
	cfg.next = locations.Next
        cfg.previous = locations.Previous
        for _, loc := range locations.Results {
                fmt.Println(loc.Name)
        }
	return nil
}
