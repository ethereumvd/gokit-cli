package main

import (
	"errors"
	"fmt"
	"log"
)

func callbackMap(c *Config, args ...string) error {

	resp, err := c.pokeapiclient.ListLocationAreas(c.nextLocationURL)
	if err != nil {
	}

	fmt.Println("Location Areas :")
	for _, area := range resp.Results {
		fmt.Println("----", area.Name)
	}

	c.nextLocationURL = resp.Next
	c.prevLocationURL = resp.Previous

	return nil
}

func callbackMapb(c *Config, args ...string) error {
	if c.prevLocationURL == nil {
		return errors.New("you are on the first page !")
	}

	resp, err := c.pokeapiclient.ListLocationAreas(c.nextLocationURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location Areas :")
	for _, area := range resp.Results {
		fmt.Println("----", area.Name)
	}

	c.nextLocationURL = resp.Next
	c.prevLocationURL = resp.Previous

	return nil
}
