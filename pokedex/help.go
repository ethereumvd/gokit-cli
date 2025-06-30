package main

import "fmt"

func callbackHelp(c *Config, args ...string) error {
	fmt.Println("Welcome to the pokedex cli menu")
	fmt.Println("here are your options")

	availCom := getCommands()
	for _, cmd := range availCom {
		fmt.Println(" --", cmd.name, ":", cmd.description)
	}

	return nil
}
