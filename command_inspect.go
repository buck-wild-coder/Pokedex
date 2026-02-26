package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	pokemonName := args[1]
	pokemon, exits := cfg.caught[pokemonName]
	if !exits {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats:")
	for _, item := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", item.Stat.Name, item.BaseStat)
	}
	fmt.Println("Types:")
	for _, item := range pokemon.Types {
		fmt.Printf("  - %s\n", item.Type.Name)
	}
	return nil
}
