package service

import (
	"sort"

	"main.go/model"
	"main.go/repo"
)

type FlightService struct {
	Repos []repo.FlightRepository
}

func (s *FlightService) GetAllFlights(sortBy string) ([]model.Flight, error) {
	var flights []model.Flight
	for _, r := range s.Repos {
		f, err := r.GetFlights()
		if err != nil {
			return nil, err
		}
		flights = append(flights, f...)
	}

	switch sortBy {
	case "price":
		sort.Slice(flights, func(i, j int) bool {
			return flights[i].Price < flights[j].Price
		})
	case "departure", "departure_time":
		sort.Slice(flights, func(i, j int) bool {
			return flights[i].DepartureTime.Before(flights[j].DepartureTime)
		})
	case "arrival", "arrival_time":
		sort.Slice(flights, func(i, j int) bool {
			return flights[i].ArrivalTime.Before(flights[j].ArrivalTime)
		})
	default:
		sort.Slice(flights, func(i, j int) bool {
			return flights[i].Price < flights[j].Price
		})
	}

	return flights, nil
}
