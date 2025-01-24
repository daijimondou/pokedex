package main

import (
	"fmt"
)

func CommandMap(cfg *Config, name string) error {
	loc, err := cfg.pokeapiClient.GetLocations(cfg.NextURL)
	if err != nil {
		return err
	}
	cfg.PreviousURL = loc.PrevURL
	cfg.NextURL = loc.NextURL
	for _, r := range loc.Results {
		fmt.Println(r.Name)
	}
	return nil
}

func CommandMapb(cfg *Config, name string) error {
	if cfg.PreviousURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	loc, err := cfg.pokeapiClient.GetLocations(cfg.PreviousURL)
	if err != nil {
		return err
	}
	cfg.PreviousURL = loc.PrevURL
	cfg.NextURL = loc.NextURL
	for _, r := range loc.Results {
		fmt.Println(r.Name)
	}
	return nil
}
