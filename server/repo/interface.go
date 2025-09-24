package repo

import "main.go/model"

type FlightRepository interface {
	GetFlights() ([]model.Flight, error)
}
