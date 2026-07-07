package commands

import (
	"fmt"

	"github.com/osamaNazieh/pokedexcli/internal/clients"
)

func listLocations(locationList []clients.LocationData) {
	for _, locaiton := range locationList {
		fmt.Println(locaiton.LocationName)
	}
}

var areaLocationsUrl = "https://pokeapi.co/api/v2/location-area"

var ConfigurationData clients.Config

func mapCommand(string) error {

	// If this is the first time we calling the list, then we want the first response with the offset of 20
	// we can check if it's the first time throught previouse were it would be and empty string
	if ConfigurationData.Previous == "" && ConfigurationData.Next == "" {
		locationList := clients.GetLocations(&ConfigurationData, areaLocationsUrl).Results
		listLocations(locationList)
		return nil
	}
	// Get the list of location from the "next" url from the response body
	locationList := clients.GetLocations(&ConfigurationData, ConfigurationData.Next)
	listLocations(locationList.Results)
	return nil
}

// get the previouse 20 locations
func mapBackCommand(string) error {
	if ConfigurationData.Previous == "" {
		fmt.Println("There isn't any previous locations to display")
		return nil
	}

	locationList := clients.GetLocations(&ConfigurationData, ConfigurationData.Previous)
	listLocations(locationList.Results)
	return nil
}
