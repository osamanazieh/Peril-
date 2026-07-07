package commands

import (
	"fmt"
	"os"
)

func exitCommand(string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil 
}