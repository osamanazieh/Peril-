package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

type config struct {
	Next string `json:"next"`
	Previous string `json:"previous"`
}

type locations struct {
	Results []locationData `json:"results"`
}

type locationData struct {
	LocationName string `json:"name"`
	Url string `json:"url"`
}

// get the byte data 
func getMarshalledResponse(urlString string) []byte {
	res, err := http.Get(urlString)
	
	if err != nil { log.Fatalf("some error happened when requestion: %v", err)}

	if res.StatusCode > 299 {
		log.Fatalf("something went wronge, not successful response: %d", res.StatusCode)
	}
	
	defer res.Body.Close()	
	binaryConfigData, err := io.ReadAll(res.Body)
	if err != nil { 
		log.Fatalf("something went wronge, when reading the body: %d", err)
	}
	return binaryConfigData
}



// get the data with the response 
func getLocationData(configurationData *config, areaLocationsUrl string) locations {
	
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


func cleanInput (text string) []string {
	result := []string{}
	
	if strings.Trim(text, " ") == "" {
		return []string{"pokedex > "}
	}

	text  = strings.Trim(text, " ")
	result = strings.Split(text, " ")	
	return result
}