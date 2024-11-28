package main

import (
	"fmt"
)

//urls of locations
type config struct {
	prevLocation string,
	currLocation string,
	nextLocation string,
}

func getLocation() string {
	res, err := http.Get("https://pokeapi.co/api/v2/location/{id or name}/")
	
}
