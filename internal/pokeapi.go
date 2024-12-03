package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/thom151/pokedexcli/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
	pokemons   []string
}

func NewClient(interval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(interval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		pokemons: []string{},
	}
}

//urls of locations

func (c *Client) GetLocation(next *string) (LocationNavigator, error) {

	locationsUrl := "https://pokeapi.co/api/v2/location-area"
	if next != nil {
		locationsUrl = *next
	}

	val, ok := c.cache.Get(locationsUrl)
	if ok {
		fmt.Print("PRINTING FROM CACHE\n")
		var cachedLocationArea LocationNavigator
		err := json.Unmarshal(val, &cachedLocationArea)
		if err != nil {
			return LocationNavigator{}, err
		}

		return cachedLocationArea, nil
	}

	req, err := http.NewRequest("GET", locationsUrl, nil)
	if err != nil {
		return LocationNavigator{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationNavigator{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationNavigator{}, err
	}

	var currentLocation LocationNavigator
	err = json.Unmarshal(body, &currentLocation)
	if err != nil {
		return LocationNavigator{}, err
	}

	c.cache.Add(locationsUrl, body)
	return currentLocation, nil
}

func (c *Client) GetPokemonFromArea(area string) (PokemonInLocation, error) {

	if len(area) < 1 {
		return PokemonInLocation{}, fmt.Errorf("no area passed")
	}

	fullUrl := "https://pokeapi.co/api/v2/location-area/" + area

	val, ok := c.cache.Get(fullUrl)
	if ok {
		fmt.Print("PRINTING FROM CACHE\n")
		var pokemonStruct PokemonInLocation
		err := json.Unmarshal(val, &pokemonStruct)
		if err != nil {
			return PokemonInLocation{}, err
		}

		return pokemonStruct, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return PokemonInLocation{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInLocation{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonInLocation{}, err
	}

	var pokemonStruct PokemonInLocation
	err = json.Unmarshal(body, &pokemonStruct)
	if err != nil {
		return PokemonInLocation{}, err
	}

	c.cache.Add(fullUrl, body)
	return pokemonStruct, nil

}

func (c *Client) GetPokemon(pokemon string) (PokemonInfo, error) {
	fullUrl := "https://pokeapi.co/api/v2/pokemon/" + pokemon

	val, ok := c.cache.Get(fullUrl)
	if ok {
		fmt.Print("PRINTING FROM CACHE\n")
		var pokemonInfo PokemonInfo
		err := json.Unmarshal(val, &pokemonInfo)
		if err != nil {
			return PokemonInfo{}, err
		}

		return pokemonInfo, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return PokemonInfo{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfo{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonInfo{}, err
	}

	var pokemonInfo PokemonInfo
	err = json.Unmarshal(body, &pokemonInfo)
	if err != nil {
		return PokemonInfo{}, nil
	}

	c.cache.Add(fullUrl, body)

	return pokemonInfo, nil

}
