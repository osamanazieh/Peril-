package commands

import "fmt"

func pokedexCommand(string) error {
	fmt.Println("Your Pokedex:")
	for name := range catchedPokemons {
		fmt.Printf("- %s\n", name)
	}
	return  nil
}