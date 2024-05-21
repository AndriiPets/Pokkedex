package main

import (
	"fmt"
	"math/rand"
	"os"

	pokeapi "github.com/AndriiPets/pokdex/internal/PokeAPI"
)

func commandExit(c *Config, args ...string) error {
	fmt.Println()
	fmt.Println("Bye, bye")
	os.Exit(0)

	return nil
}

func commandHelp(c *Config, args ...string) error {
	fmt.Println()
	fmt.Println("Wellcome to the Pokkdex!")
	fmt.Println("Usage:")
	fmt.Println()
	for name, com := range getCommands() {
		fmt.Printf("\n%s: %s", name, com.desc)
	}
	fmt.Println()

	return nil
}

func commandMap(c *Config, args ...string) error {
	if len(args) > 0 {
		fmt.Println("This command does not take any arguments\nTo get the list of avaliable commands use: help")
		return nil
	}

	locs, err := c.Client.GetLocations(c.Next)
	if err != nil {
		return err
	}
	c.Next = locs.Next
	c.Previous = locs.Previous

	for _, loc := range locs.Results {
		fmt.Printf("\n%s", loc.Name)
	}
	fmt.Println()
	return nil
}

func commandMapB(c *Config, args ...string) error {
	if len(args) > 0 {
		fmt.Println("This command does not take any arguments\nTo get the list of avaliable commands use: help")
		return nil
	}

	if prev := c.Previous; prev == nil {
		fmt.Println("You reached the end of the list")
		return nil
	}
	locs, err := c.Client.GetLocations(c.Previous)
	if err != nil {
		return err
	}
	c.Next = locs.Next
	c.Previous = locs.Previous

	for _, loc := range locs.Results {
		fmt.Printf("\n%s", loc.Name)
	}
	fmt.Println()
	return nil
}

func commandExplore(c *Config, args ...string) error {
	if len(args) <= 0 {
		fmt.Println("This command requaiers area argument\nTo get the list of avaliable commands use: help")
		return nil
	}
	area := args[0]

	fmt.Println()
	fmt.Printf("Exploring %s...", area)

	loc, err := c.Client.GetEncounters(area)
	if err != nil {
		if err == pokeapi.AreaNotFound {
			fmt.Printf("\n'%s' Area not found!", area)
			return nil
		} else {
			return err
		}
	}

	fmt.Println("\nFound Pokemon:")

	for _, encounter := range loc.PokemonEncounters {
		fmt.Printf("\n-%s", encounter.Pokemon.Name)
	}
	fmt.Println()

	return nil
}

func commandCatch(c *Config, args ...string) error {
	if len(args) <= 0 {
		fmt.Println("This command requaiers pokemon argument\nTo get the list of avaliable commands use: help")
		return nil
	}
	pokemon := args[0]

	pok, err := c.Client.GetPokemon(pokemon)
	if err != nil {
		if err == pokeapi.PokemonNotFound {
			fmt.Printf("\n'%s' Pokemon not found!", pokemon)
			return nil
		}
	}

	fmt.Println()
	fmt.Printf("Throwing a pokeball at %s...", pok.Name)

	chance := rand.Intn(pok.BaseExperience)
	if chance <= 40 {
		fmt.Printf("\n%s was caught!", pok.Name)
		fmt.Println("\nYou may now inspect it with the inspect command.")
		c.Pokedex[pok.Name] = pok
	} else {
		fmt.Printf("\n%s escaped!", pok.Name)
		return nil
	}

	return nil
}

func commandInspect(c *Config, args ...string) error {
	if len(args) <= 0 {
		if len(c.Pokedex) <= 0 {
			fmt.Printf("You havent caught any pokemon yet!")
			return nil
		}
		fmt.Println("\nYour Pokedex:")
		for _, p := range c.Pokedex {
			fmt.Printf("\n\t-%s", p.Name)
		}
		fmt.Println()
		return nil
	}
	pokemon := args[0]

	entry, ok := c.Pokedex[pokemon]
	if !ok {
		fmt.Printf("\nYou have not caught pokemon '%s'", pokemon)
		return nil
	}

	fmt.Printf("\nName: %s", entry.Name)
	fmt.Printf("\nHeight: %d", entry.Height)
	fmt.Printf("\nWeight: %d", entry.Weight)
	fmt.Println("\nStats:")
	for _, stat := range entry.Stats {
		fmt.Printf("\n\t-%s: %d", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("\nTypes:")
	for _, t := range entry.Types {
		fmt.Printf("\n\t-%s", t.Type.Name)
	}

	return nil
}
