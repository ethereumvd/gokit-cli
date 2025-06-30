package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(c *Config, args ...string) error {

	if len(args) != 1 {
		return errors.New("No Pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, err := c.pokeapiclient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	Exp := rand.Intn(pokemon.BaseExperience)

	if Exp > threshold {
		return fmt.Errorf("Failed to catch %s !", pokemonName)
	}

	c.caughtPokemons[pokemonName] = pokemon
	fmt.Printf("%s was caught ! \n", pokemonName)

	return nil
}
