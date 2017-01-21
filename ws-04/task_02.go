package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Parkings defines parking list.
type Parkings struct {
	Items []Parking `json:"Items"`
}

// Parkings defines parking place.
type Parking struct {
	Address     string   `json:"Address"`
	FreePlaces  int      `json:"FreePlaces"`
	ID          string   `json:"Id"`
	IsLocked    bool     `json:"IsLocked"`
	Position    Position `json:"Position"`
	TotalPlaces int      `json:"TotalPlaces"`
}

// Position defines parking coordinates.
type Position struct {
	Lat float64 `json:"Lat"`
	Lon float64 `json:"Lon"`
}

func main() {
	url := "http://apivelobike.velobike.ru/ride/parkings"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Couldn't get: %v\n", err)
	}

	parkings := new(Parkings)

	err = json.NewDecoder(resp.Body).Decode(parkings)

	if err != nil {
		log.Fatalf("Couldn't recognize body: %v\n", err)
	}

	fmt.Printf("Parking list is: %v\n\n\n", parkings)

	fmt.Printf("First parking is: %s, %s\n", parkings.Items[0].ID, parkings.Items[0].Address)
}
