package repo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Server2Repo struct {
	BaseURL string
}

func (r *Server2Repo) GetFlights() ([]Flight, error) {
	url := r.BaseURL + "/flights"
	fmt.Println("Fetching flights from Server1Repo:", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching flights: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var flights []Flight
	if err := json.NewDecoder(resp.Body).Decode(&flights); err != nil {
		return nil, fmt.Errorf("error decoding flights: %w", err)
	}
	return flights, nil
}
