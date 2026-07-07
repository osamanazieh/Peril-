package commands

import (
	"fmt"

	"github.com/osamaNazieh/pokedexcli/internal/clients"
)


func printPokemonTypes(pokemonTypes []clients.PokemonType) {
	fmt.Println("Types:")
	for _, t := range pokemonTypes {
		fmt.Printf("- %s\n", t.Type.Name)
	}
}

func printPokemonStats(pokemonStats []clients.PokemonStat) {
	fmt.Println("Stats:")
	for _, stat := range pokemonStats {
		fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
}




func inspectCommand(pokemonName string) error {
	pokemon, ok := catchedPokemons[pokemonName]
	if !ok {
		return fmt.Errorf("The pokemon is not captuered...")	
	}

	fmt.Printf("Name: %s\n", pokemonName)
	fmt.Printf("Height: %d\n", pokemon.ownedPokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.ownedPokemon.Weight)
	printPokemonStats(pokemon.ownedPokemon.Stats)
	printPokemonTypes(pokemon.ownedPokemon.Types)
	return nil
}