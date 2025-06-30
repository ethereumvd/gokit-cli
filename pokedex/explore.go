package main

import (
	"errors"
	"fmt"
)

func callbackExplore(c *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No location area provided")
	}
	locationAreaName := args[0]

	resp, err := c.pokeapiclient.ListLocationArea(locationAreaName)
	if err != nil {
	}

	fmt.Printf("Pokemons in %s: \n", resp.Name)
	for _, area := range resp.PokemonEncounters {
		fmt.Println("----", area.Pokemon.Name)
	}

	return nil
}
