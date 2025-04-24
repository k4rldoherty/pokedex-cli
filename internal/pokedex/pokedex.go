package pokedex

import (
	"github.com/k4rldoherty/pokedex-cli/internal/api"
)

type Pokedex struct {
	CaughtPokemon map[string]api.Pokemon
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		CaughtPokemon: make(map[string]api.Pokemon),
	}
}
