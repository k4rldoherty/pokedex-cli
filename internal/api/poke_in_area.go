package api

type PokeInArea struct {
	Results []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon PokemonInArea `json:"pokemon"`
}

type PokemonInArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
