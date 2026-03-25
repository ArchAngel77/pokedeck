package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/ArchAngel77/pokedeck/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	next 	      *string
	previous      *string
	Pokedex       map[string]pokeapi.Pokemon
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:		"map",
			description:	"Displays names and locations",
			callback:	commandMap,
		},
		"mapb": {
                        name:           "mapb",
                        description:    "Displays names and locations of previous selection",
                        callback:       commandMapBack,
                },
		 "explore": {
                        name:           "explore",
                        description:    "Explore a location area",
                        callback:       commandExplore,
                },
		  "catch": {
                        name:           "catch",
                        description:    "Catch a Pokemon",
                        callback:       commandCatch,
                },
		"inspect": {
                        name:           "inspect",
                        description:    "Inspect a caught Pokemon",
                        callback:       commandInspect,
                },
		"pokedex": {
                        name:           "pokedex",
                        description:    "List caught Pokemon",
                        callback:       commandPokedex,
                },
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	result := strings.Fields(text)
	return result
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		args := words[1:]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
