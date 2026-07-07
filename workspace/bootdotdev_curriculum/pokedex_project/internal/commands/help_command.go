package commands

import (
	"fmt"
)

func helpCommand(string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")	
	for _, val := range Commands {
		fmt.Printf("%s: %s\n", val.Name, val.Description)
	}
	return nil
}