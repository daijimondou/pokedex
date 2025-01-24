package main

import "fmt"

func CommandPokedex(cfg *Config, name string) error {
	fmt.Println("Your Pokedex:")
	for k := range cfg.Pokedex {
		fmt.Printf("- %v\n", k)
	}
	return nil
}
