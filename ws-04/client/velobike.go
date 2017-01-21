package velobike

import (
	"encoding/json"
	"net/http"
)

type VelobikeClient struct {
	baseURL string
	key string

	ParkingService *ParkingService
}

func NewVelobikeClient(baseURL, key string) *VelobikeClient {
	client := &VelobikeClient{
		baseURL: baseURL,
	}
	return client
}

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


func (v *VelobikeClient) GetParkings() (*Parkings, error) {
	url := v.baseURL + "/ride/parkings?key=" + v.key

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	parkings := new(Parkings)

	err = json.NewDecoder(resp.Body).Decode(parkings)

	if err != nil {
		return nil, err
	}

	return parkings, nil
}
