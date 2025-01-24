package main

import (
	"fmt"
)

func CommandInspect(cfg *Config, name string) error {
	poke, ok := cfg.Pokedex[name]
	if !ok {
		return fmt.Errorf("%v is not caught yet", name)
	}
	fmt.Printf("Name: %v\n", name)
	fmt.Printf("Height: %v\n", poke.Height)
	fmt.Printf("Weight: %v\n", poke.Weight)
	fmt.Println("Stats:")
	for _, v := range poke.Stats {
		fmt.Printf("-%v: %v\n", v.Name, v.Value)
	}
	fmt.Println("Types:")
	for _, t := range poke.Types {
		fmt.Printf("- %v\n", t)
	}
	return nil
}
