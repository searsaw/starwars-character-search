package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const swapiEndpoint = "https://swapi.co/api"

type (
	swapiCharacter struct {
		Name      string    `json:"name"`
		Height    string    `json:"height"`
		Mass      string    `json:"mass"`
		HairColor string    `json:"hair_color"`
		SkinColor string    `json:"skin_color"`
		EyeColor  string    `json:"eye_color"`
		BirthYear string    `json:"birth_year"`
		Gender    string    `json:"gender"`
		Homeworld string    `json:"homeworld"`
		Films     []string  `json:"films"`
		Species   []string  `json:"species"`
		Vehicles  []string  `json:"vehicles"`
		Starships []string  `json:"starships"`
		Created   time.Time `json:"created"`
		Edited    time.Time `json:"edited"`
		URL       string    `json:"url"`
	}

	swapiResult struct {
		Count    int              `json:"count"`
		Next     string           `json:"next"`
		Previous string           `json:"previous"`
		Results  []swapiCharacter `json:"results"`
	}
)

func main() {
	term := os.Args[1]
	url := swapiEndpoint + "/people?search=" + term

	resp, err := http.Get(url)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	defer resp.Body.Close()

	var sr swapiResult
	err = json.NewDecoder(resp.Body).Decode(&sr)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Printf("----- Star Wars Characters Searched By '%s' -----\n", term)
	for _, character := range sr.Results {
		fmt.Println("Character name:", character.Name)
	}
}
