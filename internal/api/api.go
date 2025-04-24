package api

import (
	"encoding/json"
	"io"
)

const URL = "https://pokeapi.co/api/v2/"

func (c *Client) GetPokemonAreas(pageUrl *string) (PokemonAreaResults, error) {
	resourceUrl := URL + "location-area"
	if pageUrl != nil {
		resourceUrl = *pageUrl
	}

	var areas PokemonAreaResults
	var data []byte

	// check the cache before making request
	if data, ok := c.cache.Get(resourceUrl); ok {
		err := json.Unmarshal(data, &areas)
		if err != nil {
			return PokemonAreaResults{}, nil
		}
		return areas, nil
	}

	res, err := c.httpClient.Get(resourceUrl)
	if err != nil {
		return PokemonAreaResults{}, err
	}
	defer res.Body.Close()
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return PokemonAreaResults{}, err
	}
	// otherwise add the data to the cache
	c.cache.Add(resourceUrl, data)
	if err = json.Unmarshal(data, &areas); err != nil {
		return PokemonAreaResults{}, err
	}
	return areas, nil
}

func (c *Client) GetPokemonInArea(area string) (PokeInArea, error) {
	resourceUrl := URL + "location-area/" + area
	var pokesInArea PokeInArea
	var data []byte
	if data, ok := c.cache.Get(resourceUrl); ok {
		err := json.Unmarshal(data, &pokesInArea)
		if err != nil {
			return PokeInArea{}, err
		}
		return pokesInArea, nil
	}
	res, err := c.httpClient.Get(resourceUrl)
	if err != nil {
		return PokeInArea{}, err
	}
	defer res.Body.Close()
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return PokeInArea{}, err
	}
	c.cache.Add(resourceUrl, data)
	if err = json.Unmarshal(data, &pokesInArea); err != nil {
		return PokeInArea{}, err
	}
	return pokesInArea, nil
}

func (c *Client) GetPokemonByName(pokemon string) (Pokemon, error) {
	resourceUrl := URL + "pokemon/" + pokemon
	var pokemonData Pokemon
	var data []byte
	if data, ok := c.cache.Get(resourceUrl); ok {
		err := json.Unmarshal(data, &pokemonData)
		if err != nil {
			return Pokemon{}, nil
		}
	}
	res, err := c.httpClient.Get(resourceUrl)
	if err != nil {
		return Pokemon{}, nil
	}
	defer res.Body.Close()
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(resourceUrl, data)
	if err = json.Unmarshal(data, &pokemonData); err != nil {
		return Pokemon{}, err
	}
	return pokemonData, nil
}
