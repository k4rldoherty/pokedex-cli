package api

type PokemonAreaResults struct {
	Next     *string       `json:"next"`
	Previous *string       `json:"previous"`
	Results  []PokemonArea `json:"results"`
}

type PokemonArea struct {
	Name string `json:"name"`
}
