package main


type cliCommands struct {
	name string 
	description string 
	command func() error 
}


var commands map[string] cliCommands

func init() {
	commands = map[string] cliCommands {
		"help": {
			name: "help",
			description: "show all the available commands with thier description",
			command: helpCommand,
		},
		"exit" :{ 
			name: "exit", 
			description: "end and exit the program", 
			command: exitCommand,
		},
		"map" :{ 
			name: "map", 
			description: "list 20 location, up on subsequent command will call the next 20 lcoations", 
			command: mapCommand,
		},
		"mapb" :{ 
			name: "mapb", 
			description: "list previous 20 location, up on subsequent command will call the previous 20 lcoations", 
			command: mapBackCommand,
		},
	}
}