package repl

import (
	"fmt"
	"os"

	"github.com/k4rldoherty/pokedex-cli/internal/config"
)

type cliCommand struct {
	name     string
	desc     string
	callback func(*config.Config, []string) error
}

var (
	ReplCommands map[string]cliCommand
	cfg          config.Config
)

// Prevents an init cycle where handleHelpCommand is dependent on ReplCommands.
func InitCommands() {
	ReplCommands = map[string]cliCommand{
		"exit": {
			name:     "exit",
			desc:     "Exits the function.",
			callback: exitRepl,
		},
		"help": {
			name:     "help",
			desc:     "Returns available commands from this tool.",
			callback: handleHelpCommand,
		},
		"map": {
			name:     "map",
			desc:     "Returns a map of pokemon areas",
			callback: handleMapFCommand,
		},
		"mapb": {
			name:     "mapb",
			desc:     "Returns the previous page of the pokemon areas",
			callback: handleMapBCommand,
		},
		"explore": {
			name:     "explore",
			desc:     "Shows a list of all pokemon available in this area",
			callback: handleExploreCommand,
		},
	}
}

func exitRepl(c *config.Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func handleHelpCommand(c *config.Config, args []string) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Useage:")
	for _, v := range ReplCommands {
		fmt.Printf("%v: %v\n", v.name, v.desc)
	}
	return nil
}

func handleMapFCommand(cfg *config.Config, args []string) error {
	areas, err := cfg.Client.GetPokemonAreas(cfg.Next)
	if err != nil {
		return err
	}
	cfg.Next = areas.Next
	cfg.Previous = areas.Previous
	for _, p := range areas.Results {
		fmt.Println(p.Name)
	}
	return nil
}

func handleMapBCommand(cfg *config.Config, args []string) error {
	if cfg.Previous == nil {
		fmt.Println("You are already at page 1")
		return nil
	}
	areas, err := cfg.Client.GetPokemonAreas(cfg.Previous)
	if err != nil {
		return err
	}
	cfg.Next = areas.Next
	cfg.Previous = areas.Previous
	for _, p := range areas.Results {
		fmt.Println(p.Name)
	}
	return nil
}

func handleExploreCommand(cfg *config.Config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("incorrect number of arguments passed")
	}
	area := args[0]
	fmt.Printf("Exploring %v area...\n", area)
	pokemon, err := cfg.Client.GetPokemonInArea(area)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, p := range pokemon.Results {
		fmt.Println(" - ", p.Pokemon.Name)
	}
	return nil
}
