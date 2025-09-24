package service

import (
	"sort"

	"main.go/repo"
	"main.go/utils"
)

type FlightService struct {
	Repos []repo.FlightRepository
}

func (s *FlightService) GetAllFlights(sortBy string) ([]repo.Flight, error) {
	var flights []repo.Flight
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
	case "departure":
		sort.Slice(flights, func(i, j int) bool {
			return flights[i].DepartureTime.Before(flights[j].DepartureTime)
		})
	case "arrival_time":
		sort.Slice(flights, func(i, j int) bool {
			return flights[i].ArrivalTime.Before(flights[j].ArrivalTime)
		})
	default:
		flights = utils.SortByPrice(flights)
	}

	return flights, nil
}
