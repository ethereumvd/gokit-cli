package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	name        string
	description string
	callback    func(*Config) error
}

func getCommands() map[string]Command {
	return map[string]Command{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Turns off pokedex",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "List location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List location areas one url back",
			callback:    callbackMapb,
		},
	}
}

func standardiseInp(inp string) []string {
	lowerCaseInp := strings.ToLower(inp)
	args := strings.Fields(lowerCaseInp)
	return args
}

func StartCLI(cfg *Config) {
	fmt.Println("hello")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">>> ")

		scanner.Scan()
		text := scanner.Text()
		args := standardiseInp(text)

		if len(args) == 0 {
			continue
		}

		primaryArg := args[0]

		availableCommands := getCommands()
		command, ok := availableCommands[primaryArg]

		if !ok {
			fmt.Println("Invaid command")
			fmt.Println("pls read help manual")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}

}
