package main

import (
	"pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &Config{
		Pokedex:       map[string]Pokemon{},
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
