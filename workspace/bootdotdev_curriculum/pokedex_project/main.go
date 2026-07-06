package main

import (
	"bufio"
	"os"
	"fmt"
	
)


func main() {
	
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("pokedex > ")
		for scanner.Scan() {
			inputCommand := cleanInput(scanner.Text())[0]
			if cmd, ok := commands[inputCommand]; ok {
				cmd.command()
			}
			break
		}
		if err := scanner.Err(); err != nil { 
			fmt.Printf("Invalid Input: %v", err)
		}
		
	}
}
