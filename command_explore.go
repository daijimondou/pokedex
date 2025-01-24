package main

import "fmt"

func CommandExplore(cfg *Config, name string) error {
	locD, err := cfg.pokeapiClient.GetExplore(name)
	if err != nil {
		return err
	}
	for _, poke := range locD.PokemonEncounters {
		fmt.Println(poke.Pokemon.Name)
	}
	return nil
}
