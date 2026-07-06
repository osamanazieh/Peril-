package main

import (
	"fmt"
)

func helpCommand() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")	
	for _, val := range commands {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}
	return nil
}