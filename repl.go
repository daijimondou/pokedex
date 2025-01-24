package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"strings"
)

type Config struct {
	pokeapiClient pokeapi.Client
	NextURL       *string
	PreviousURL   *string
	Pokedex       map[string]Pokemon
}

type Pokemon struct {
	Height int
	Weight int
	Stats  []Stat
	Types  []string
}

type Stat struct {
	Name  string
	Value int
}

func startRepl(cfg *Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		argument := ""
		if len(words) > 1 {
			argument = words[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, argument)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	var result []string
	text = strings.ToLower(text)
	start := 0
	length := len(text)
	for i := 0; i < length; i++ {
		if text[i] == ' ' {
			if start < i {
				result = append(result, text[start:i])
			}
			start = i + 1
		}
	}
	if start < length {
		result = append(result, text[start:length])
	}
	return result
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config, name string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    CommandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    CommandExit,
		},
		"map": {
			name:        "map",
			description: "Displays 20 locations",
			callback:    CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations",
			callback:    CommandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays pokemons of that zone",
			callback:    CommandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Tries to catch a pokemon",
			callback:    CommandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspects the stats of a pokemon",
			callback:    CommandInspect,
		},
		"c": {
			name:        "catch",
			description: "Tries to catch a pokemon",
			callback:    CommandCatch,
		},
	}
}
