package commands

import (
	"fmt"
	"math/rand"
	"github.com/osamaNazieh/pokedexcli/internal/clients"
)

var catchedPokemons = make(map[string]Pokemon)


type Pokemon struct {
	ownedPokemon clients.PokemonData

}


func catchCommand(pokemonName string) error {
	pathUrl := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName)
	if _, ok := catchedPokemons[pokemonName]; ok {
		fmt.Printf("%s is already captured\n", pokemonName)
		return nil 	
		} 
	pokemonData := clients.CatchPokemon(pathUrl)
	baseExperience := pokemonData.BaseExperience
	catchChance := rand.Intn(baseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	if catchChance > (baseExperience / 2) {
		fmt.Printf("%s was caught!\n", pokemonName)
		catchedPokemons[pokemonName] = Pokemon{ ownedPokemon: pokemonData }
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	return nil
} 