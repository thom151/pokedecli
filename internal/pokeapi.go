package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

type LocationNavigator struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

//urls of locations

func (c *Client) GetLocation() (LocationNavigator, error) {

	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/location-area", nil)
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

	var currentLocationArea LocationNavigator
	err = json.Unmarshal(body, &currentLocationArea)
	if err != nil {
		return LocationNavigator{}, err
	}

	return currentLocationArea, nil
}

func CallMap() error {
	client := NewClient()
	locations, err := client.GetLocation()
	if err != nil {
		return err
	}

	for _, value := range locations.Results {
		fmt.Printf("%s\n", value.Name)
	}

	fmt.Println()
	return nil
}
