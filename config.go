package main

import (
	"time"

	pokeapi "github.com/AndriiPets/pokdex/internal/PokeAPI"
)

type cliCommand struct {
	name     string
	desc     string
	callback func(*Config, ...string) error
}

type Config struct {
	Client   pokeapi.Client
	Next     *string
	Previous *string
	Pokedex  map[string]pokeapi.Pokemon
}

func NewConfig() *Config {
	client := pokeapi.NewClient(time.Second * 5)
	conf := Config{
		Client:  client,
		Pokedex: make(map[string]pokeapi.Pokemon),
	}

	return &conf
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:     "help",
			desc:     "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name:     "exit",
			desc:     "Exits the Pokkdex",
			callback: commandExit,
		},
		"map": {
			name:     "map",
			desc:     "Get a list of locations",
			callback: commandMap,
		},
		"mapb": {
			name:     "mapb",
			desc:     "Go back to the previous list",
			callback: commandMapB,
		},
		"explore": {
			name:     "explore",
			desc:     "explore <location_name> - get list of all posible pokemon encounters in the area",
			callback: commandExplore,
		},
		"catch": {
			name:     "catch",
			desc:     "catch <pokemon_name> - attempt to catch a pokemon, attempt success is based on pokemon base exp",
			callback: commandCatch,
		},
		"inspect": {
			name:     "inspect",
			desc:     "look at all pokemon entries in pokedex\n\tinspect <pokemon_name> - look at the pokemon stats, if pokemon entry in in the pokedex",
			callback: commandInspect,
		},
	}
}
