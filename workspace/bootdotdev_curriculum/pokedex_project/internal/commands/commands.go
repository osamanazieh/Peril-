package commands


type cliCommands struct {
	Name string 
	Description string 
	Command func(string) error 
}


var Commands map[string] cliCommands

func init() {
	Commands = map[string] cliCommands {
		"help": {
			Name: "help",
			Description: "show all the available commands with thier description",
			Command: helpCommand ,
		},
		"exit" :{ 
			Name: "exit", 
			Description: "end and exit the program", 
			Command: exitCommand,
		},
		"map" :{ 
			Name: "map", 
			Description: "list 20 location, up on subsequent command will call the next 20 lcoations", 
			Command: mapCommand,
		},
		"mapb" :{ 
			Name: "mapb", 
			Description: "list previous 20 location, up on subsequent command will call the previous 20 lcoations", 
			Command: mapBackCommand,
		},
		"explore" :{ 
			Name: "explore", 
			Description: "see what pokemon in this area, this command need the nsame of the area as the second arrgument", 
			Command: exploreCommand,
		},
		"catch" :{ 
			Name: "catch", 
			Description: "cathc a pokemon, this command need the name of the pokemon", 
			Command: catchCommand,
		},
		"inspect" :{ 
			Name: "inspect", 
			Description: "Display the information about specific captuered pokemon, the pokemon name must be provided with this command", 
			Command: inspectCommand,
		},
		"pokedex" :{ 
			Name: "pokedex", 
			Description: "show what pokemons you have in your pokedex", 
			Command: pokedexCommand,
		},
	}
}