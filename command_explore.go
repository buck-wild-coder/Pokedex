package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) < 2 {
		return errors.New("Please add the name/id of location you want to explore")
	}

	listPokemon, err := cfg.pokeapiClient.ListPokemon(args[1])
	if err != nil {
		return err
	}

	for _, loc := range listPokemon.PokemonEncounters {
		fmt.Println(loc.Pokemon.Name)
	}
	return nil
}
