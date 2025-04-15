package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/k4rldoherty/pokedex-cli/internal/config"
)

func StartRepl(cfg *config.Config) {
	// initializes the map of commands
	InitCommands()
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		commands := CleanInput(reader.Text())
		if len(commands) == 0 {
			continue
		}
		command, ok := ReplCommands[commands[0]]
		if !ok {
			fmt.Println("Unknown command.")
			continue
		}
		// calls the function associated with the command, passing in the config of the application
		if e := command.callback(cfg, commands[1:]); e != nil {
			fmt.Println("Error: ", e)
		}
		fmt.Println()
	}
}

// Parses input into a slice of strings and converts to lower case
func CleanInput(s string) []string {
	return strings.Fields(strings.ToLower(s))
}
