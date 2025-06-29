package main

import "github.com/ethereumvd/gokit-cli/pokedex/internal/pokeapi"

type Config struct {
	pokeapiclient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
}

func main() {
	cfg := Config{
		pokeapiclient: pokeapi.NewClient(),
	}

	StartCLI(&cfg)

}
