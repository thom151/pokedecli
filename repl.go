package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			if command, ok := getCommands()[input]; ok {
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
			callback:    callMap,
		},
		"mapb": {
			name:        "mapb",
			description: "shows previous 20 poke location",
			callback:    callMapB,
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
