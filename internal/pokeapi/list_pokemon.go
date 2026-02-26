package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListPokemon(name string) (PokemonLocation, error) {
	url := baseURL + "/location-area/" + name

	if val, ok := c.cache.Get(url); ok {
		locationsResp := PokemonLocation{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return PokemonLocation{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonLocation{}, nil
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonLocation{}, nil
	}

	pokemonlist := PokemonLocation{}
	err = json.Unmarshal(dat, &pokemonlist)
	if err != nil {
		return PokemonLocation{}, nil
	}

	c.cache.Add(url, dat)
	return pokemonlist, nil
}
