package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) Pokemoninfo(pageURL *string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + *pageURL

	if val, ok := c.cache.Get(*pageURL); ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.Get(url)
	if err != nil {
		return Pokemon{}, err
	}

	dat, err := io.ReadAll(req.Body)
	if err != nil {
		return Pokemon{}, nil
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(dat, &pokemon)
	if err != nil {
		return Pokemon{}, nil
	}
	c.cache.Add(*pageURL, dat)

	return pokemon, nil
}
