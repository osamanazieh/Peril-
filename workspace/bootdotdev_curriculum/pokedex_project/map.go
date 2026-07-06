package main

import (
	"fmt"
)


func listLocations(locationList []locationData) {
	for _, locaiton := range locationList {
		fmt.Println(locaiton.LocationName)
	}
}

var areaLocationsUrl = "https://pokeapi.co/api/v2/location-area"

var ConfigurationData config 

func mapCommand()  error {

	// If this is the first time we calling the list, then we want the first response with the offset of 20  
	// we can check if it's the first time throught previouse were it would be and empty string 
	if ConfigurationData.Previous == "" && ConfigurationData.Next == "" {
		locationList := getLocationData(&ConfigurationData, areaLocationsUrl).Results
		listLocations(locationList)
		return nil
	}
	// Get the list of location from the "next" url from the response body 
	locationList := getLocationData(&ConfigurationData, ConfigurationData.Next)
	listLocations(locationList.Results)
	return nil 
}

// get the previouse 20 locations
func mapBackCommand() error {
	if ConfigurationData.Previous == "" {
		fmt.Println("There isn't any previous locations to display")
		return nil 
	}

	locationList := getLocationData(&ConfigurationData, ConfigurationData.Previous)
	listLocations(locationList.Results)
	return nil
}