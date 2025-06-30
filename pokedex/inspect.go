package main

import (
	"errors"
	"fmt"
)

func callbackInspect(c *Config, args ...string) error {

	if len(args) != 1 {
		return errors.New("No Pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, ok := c.caughtPokemons[pokemonName]
	if !ok {
		return fmt.Errorf("You havent caught %s yet !", pokemonName)
	}

	fmt.Printf("Name of Pokemon : %s\n", pokemon.Name)
	fmt.Printf("Height : %d \n", pokemon.Height)
	fmt.Printf("Weight : %d \n", pokemon.Weight)

	fmt.Printf("Abilities :\n")
	for _, ab := range pokemon.Abilities {
		fmt.Printf("---- %s\n", ab.Ability.Name)
	}

	fmt.Printf("Stats :\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("---- %s : %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Base Exp of Pokemon : %d\n", pokemon.BaseExperience)

	return nil
}
