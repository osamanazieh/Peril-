package main

import (
	"fmt"
	"os"
)

func exitCommand() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil 
}