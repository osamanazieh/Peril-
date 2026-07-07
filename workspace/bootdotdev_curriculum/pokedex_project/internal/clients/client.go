package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"github.com/osamaNazieh/pokedexcli/internal/pokecache"
)

// Get a location List, used for "map" and "mapb" command 
type Config struct {
	Next string `json:"next"`
	Previous string `json:"previous"`
}

type locations struct {
	Results []LocationData `json:"results"`
}

type LocationData struct {
	LocationName string `json:"name"`
	Url string `json:"url"`
}

// get the all Pokemons in some area, used "explore" command
type locationArea struct {
	PokemonEncounters []pokemonEncounters `json:"pokemon_encounters"` 
}
type pokemonEncounters struct {
	Pokemons Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
}

type PokemonData struct {
    Height         int           `json:"height"`
    Weight         int           `json:"weight"`
    BaseExperience int           `json:"base_experience"`
    Types          []PokemonType `json:"types"`
    Stats          []PokemonStat `json:"stats"`
}

type PokemonType struct {
    Type Type `json:"type"`
}

type Type struct {
    Name string `json:"name"`
}

type PokemonStat struct {
    BaseStat int  `json:"base_stat"`
    Stat     Stat `json:"stat"`
}

type Stat struct {
    Name string `json:"name"`
}

var cashedData = pokecache.NewCahse(5 * time.Second)


// get the byte data 
func getMarshalledResponse(urlString string) []byte {
	if data, ok := cashedData.Get(urlString); ok {
		return data
	}
	res, err := http.Get(urlString)
	
	if err != nil { log.Fatalf("some error happened when requestion: %v", err)}
	
	if res.StatusCode > 299 {
		log.Fatalf("something went wronge, not successful response: %d", res.StatusCode)
	}
	
	defer res.Body.Close()	
	dataAsBytes, err := io.ReadAll(res.Body)
	cashedData.Add(urlString, dataAsBytes)
	
	if err != nil { 
		log.Fatalf("something went wronge, when reading the body: %d", err)
	}
	return dataAsBytes
}



// get the data with the response 
func GetLocations(configurationData *Config, areaLocationsUrl string) locations {
	
	dataAsBytes := getMarshalledResponse(areaLocationsUrl)
	var results locations
	if err := json.Unmarshal(dataAsBytes, &results); err != nil {
		log.Fatalf("something went wronge, when unmarshalling the data: %v", err)
	}
	
	if err := json.Unmarshal(dataAsBytes, configurationData); err != nil {
		log.Fatalf("something went wronge, when unmarshalling the config data: %v", err)
	}
	
	return results
} 


func GetLocationPokemonEncounter(areaName string) []pokemonEncounters {
	pathurl := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", areaName)
	bytesData := getMarshalledResponse(pathurl)	

	var pokemonEncouters locationArea
	if err := json.Unmarshal(bytesData, &pokemonEncouters); err != nil {
		log.Fatal(err)
	}
	return pokemonEncouters.PokemonEncounters
}


func CatchPokemon(pathUrl string) PokemonData {
	bytesAsData := getMarshalledResponse(pathUrl)

	var pokemonData PokemonData
	if err := json.Unmarshal(bytesAsData, &pokemonData); err != nil {
		fmt.Errorf("something is wrong with the unmarshalling the context: %s", err)
	}
	// fmt.Println(pokemonData)
	return pokemonData
}