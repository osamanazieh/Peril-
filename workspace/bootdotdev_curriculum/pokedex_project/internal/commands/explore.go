package commands

import (
	"fmt"

	"github.com/osamaNazieh/pokedexcli/internal/clients"
)

func exploreCommand(areaName string) error {
	if areaName == "" {
		return fmt.Errorf("the name of the area wasn't provided") 
	}
	pokemonEcounters := clients.GetLocationPokemonEncounter(areaName)
	fmt.Printf("Exploring %s...\nFound Pokemon:\n", areaName)
	for _, pokemonStruct := range pokemonEcounters {
		fmt.Printf("- %s\n", pokemonStruct.Pokemons.Name)
	}
	
	return nil
}