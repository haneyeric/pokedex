package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokenmon(name string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + name

	if val, ok := c.cache.Get(url); ok {
		pokeResp := RespPokemon{}
		err := json.Unmarshal(val, &pokeResp)
		if err != nil {
			return RespPokemon{}, err
		}

		return pokeResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	pokeResp := RespPokemon{}
	err = json.Unmarshal(dat, &pokeResp)
	if err != nil {
		return RespPokemon{}, err
	}

	c.cache.Add(url, dat)
	return pokeResp, nil
}
