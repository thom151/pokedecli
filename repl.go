package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func startRepl(c *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			cleaned := cleanInput(input)
			arg := ""

			if len(cleaned) > 1 {
				arg = cleaned[1]
			}

			if command, ok := getCommands()[cleaned[0]]; ok {

				err := command.callback(c, arg)
				if err != nil {
					fmt.Println(err)
				}
				continue
			} else {
				fmt.Println("Unknow command")
				continue
			}
		}
	}
}

func cleanInput(str string) []string {
	lower := strings.ToLower(str)
	words := strings.Fields(lower)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config, arg string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{

		"help": {
			name:        "help",
			description: "displays a help message",
			callback:    callHelp,
		},
		"exit": {
			name:        "exit",
			description: "exists the program",
			callback:    callExit,
		},
		"map": {
			name:        "map",
			description: "shows next 20 poke location",
			callback:    CallMap,
		},
		"mapb": {
			name:        "mapb",
			description: "shows previous 20 poke location",
			callback:    CallMapB,
		},
		"explore": {
			name:        "explore",
			description: "see pokemons in a given area",
			callback:    CallExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try and catch pokemons",
			callback:    CallCatch,
		},
		"inspect": {
			name:        "insepct",
			description: "see pokemon details",
			callback:    CallInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "list all pokemon caughts",
			callback:    CallPokedex,
		},
	}
}

//COMMANDS

func callHelp(c *config, arg string) error {
	fmt.Println()
	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()
	commands := getCommands()
	for _, value := range commands {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}

	fmt.Println()
	return nil
}

func callExit(c *config, arg string) error {
	os.Exit(0)
	return nil
}

func CallMap(c *config, arg string) error {
	locations, err := c.client.GetLocation(c.nextLocation)
	if err != nil {
		return err
	}

	fmt.Println()
	for _, value := range locations.Results {
		fmt.Printf("%s\n", value.Name)
	}

	c.nextLocation = locations.Next
	c.prevLocation = locations.Previous
	fmt.Println()
	return nil
}

func CallMapB(c *config, arg string) error {
	if c.prevLocation == nil {
		return fmt.Errorf("No previous locations")
	}
	locations, err := c.client.GetLocation(c.prevLocation)

	if err != nil {
		return err
	}

	fmt.Println()
	for _, value := range locations.Results {
		fmt.Printf("%s\n", value.Name)
	}

	c.nextLocation = locations.Next
	c.prevLocation = locations.Previous

	return nil

}

func CallExplore(c *config, arg string) error {

	if len(arg) < 1 {
		return fmt.Errorf("Location parameter missing. Please provide one")
	}

	encounters, err := c.client.GetPokemonFromArea(arg)

	if err != nil {
		return err
	}

	for _, pokemon := range encounters.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	fmt.Println()
	return nil
}

func CallCatch(c *config, pokemon string) error {
	if len(pokemon) < 1 {
		return fmt.Errorf("No pokemon name provided. Please provide one\n")
	}

	pokemonInfo, err := c.client.GetPokemon(pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	experience := rand.Intn(pokemonInfo.BaseExperience)
	if experience > 50 {
		return fmt.Errorf("Failed to catch %s\n", pokemon)
	}

	fmt.Printf("%s was caught!\n", pokemon)
	c.pokemons[pokemonInfo.Name] = pokemonInfo
	return nil
}

func CallInspect(c *config, pokemon string) error {
	if len(pokemon) < 1 {
		return fmt.Errorf("No pokemon name provided, please fucking provide one\n")
	}

	value, ok := c.pokemons[pokemon]
	if !ok {
		return fmt.Errorf("You ambitious fuck. You haven't caught this bro\n")
	}

	fmt.Printf("Name: %s\n", value.Name)
	fmt.Printf("Height: %d\n", value.Height)
	fmt.Printf("Weight: %d\n", value.Weight)
	fmt.Printf("Stats: \n")
	for _, stat := range value.Stats {
		fmt.Printf("\t-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types: \n")
	for _, t := range value.Types {
		fmt.Printf("\t-%s\n", t.Type.Name)
	}
	return nil

}

func CallPokedex(c *config, arg string) error {
	fmt.Println("Your pokedex: ")
	for val, _ := range c.pokemons {
		fmt.Printf("-%s", val)

	}
	fmt.Println("")
	return nil
}
