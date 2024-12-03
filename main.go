package main

import (
	"time"

	"github.com/thom151/pokedexcli/internal"
)

type config struct {
	client       internal.Client
	nextLocation *string
	prevLocation *string
	pokemons     map[string]internal.PokemonInfo
}

func main() {

	cfg := config{internal.NewClient(time.Hour), nil, nil, make(map[string]internal.PokemonInfo)}
	startRepl(&cfg)
}
