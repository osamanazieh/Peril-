package util

import (
	"strings"
)



func CleanInput (text string) []string {
	result := []string{}
	
	if strings.Trim(text, " ") == "" {
		return []string{"pokedex > "}
	}

	text  = strings.Trim(text, " ")
	result = strings.Split(text, " ")	
	return result
}