package main

import "os"

func callbackExit(c *Config, args ...string) error {
	os.Exit(0)
	return nil
}
