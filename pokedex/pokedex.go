package main

import (
	"fmt"
)

func callbackPokedex(c *Config, args ...string) error {

	if len(c.caughtPokemons) == 0 {
		return fmt.Errorf("You haven't caught any Pokemons Yet !\n Explore !")
	}

	fmt.Println("Pokemons in Pokedex :")
	for pokemon, _ := range c.caughtPokemons {
		fmt.Println(" ---", pokemon)
	}

	return nil
}