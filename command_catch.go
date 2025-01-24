package main

import (
	"fmt"
	"math/rand"
)

func CommandCatch(cfg *Config, name string) error {
	fmt.Printf("Throwing a Pokeball at %v...\n", name)
	PokeStat, err := cfg.pokeapiClient.GetStat(name)
	if err != nil {
		return err
	}
	difficulty := PokeStat.BaseExperience
	probability := 80.0
	if difficulty > 40 {
		probability -= float64(difficulty-40) * 0.132
	}
	chance := rand.Intn(99) + 1
	fmt.Printf("chance: %v vs Target: %v\n", chance, probability)
	if chance < int(probability) {
		newPoke := Pokemon{
			Height: PokeStat.Height,
			Weight: PokeStat.Weight,
		}
		for _, stat := range PokeStat.Stats {
			values := Stat{
				Name:  stat.Stat.Name,
				Value: stat.BaseStat,
			}
			newPoke.Stats = append(newPoke.Stats, values)
		}
		for _, ptype := range PokeStat.Types {
			newPoke.Types = append(newPoke.Types, ptype.Type.Name)
		}
		cfg.Pokedex[name] = newPoke
		fmt.Printf("%v was caught!\n", name)
	} else {
		fmt.Printf("failed to caught %v...\n", name)
	}

	return nil
}
