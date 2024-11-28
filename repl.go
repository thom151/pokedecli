package main

import (
	"bufio"
	"fmt"
	"github.com/thom151/pokedexcli/internal"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			cleaned := cleanInput(input)

			if command, ok := getCommands()[cleaned[0]]; ok {
				err := command.callback()
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
	callback    func() error
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
			callback:    internal.CallMap,
		},
	}
}

func callHelp() error {
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

func callExit() error {
	os.Exit(0)
	return nil
}
