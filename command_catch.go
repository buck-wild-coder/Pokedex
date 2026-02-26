package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 2 {
		return errors.New("Please insert the name of pokemon with the catch command")
	}

	pokemonName := args[1]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	pokemon, err := cfg.pokeapiClient.Pokemoninfo(&pokemonName)
	if err != nil {
		return err
	}
	if pokemon.Name == "" {
		return errors.New("Pokemon does not exist")
	}

	experience := pokemon.Base_experience
	number := rand.Intn(608)
	if experience < number {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.caught[pokemonName] = pokemon
		fmt.Println("You may now inspect it with the inspect command.")
		return nil
	}

	fmt.Printf("%s escaped!\n", pokemonName)
	return nil
}
