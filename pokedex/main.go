package main

import (
	"time"

	"github.com/ethereumvd/gokit-cli/pokedex/internal/pokeapi"
)

type Config struct {
	pokeapiclient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
	caughtPokemons  map[string]pokeapi.Pokemon
}

func main() {
	cfg := Config{
		pokeapiclient:  pokeapi.NewClient(time.Hour),
		caughtPokemons: make(map[string]pokeapi.Pokemon),
	}

	StartCLI(&cfg)
}
