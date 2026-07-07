package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/osamaNazieh/pokedexcli/internal/commands"
	"github.com/osamaNazieh/pokedexcli/internal/util"
)


func main() {
	
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("pokedex > ")
		for scanner.Scan() {
			inputCommand := util.CleanInput(scanner.Text())
			
			
			
			if cmd, ok := commands.Commands[inputCommand[0]]; ok {
				switch(cmd.Name) {
					case "explore":
						fallthrough
					case "catch":
						fallthrough
					case "inspect":
						if len(inputCommand) < 2 {
							fmt.Println("a second argument must be provided for this command, see description below")
							fmt.Printf("======================================\n%s\n======================================\n", cmd.Description)
							break
						}
						cmd.Command(inputCommand[1])
					default:
						cmd.Command("")
				}
				
			}
			break
		}
		if err := scanner.Err(); err != nil { 
			fmt.Printf("Invalid Input: %v", err)
		}
		
	}
}
