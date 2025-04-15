package main

import (
	"time"

	"github.com/k4rldoherty/pokedex-cli/internal/config"
	"github.com/k4rldoherty/pokedex-cli/internal/repl"
)

func main() {
	config := config.NewConfig(time.Minute * 5)
	repl.StartRepl(config)
}
