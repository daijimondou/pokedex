package main

import (
	"fmt"
	"os"
)

func CommandExit(cfg *Config, name string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
