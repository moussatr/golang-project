package utils

import (
	"sort"

	"main.go/model"
)

func SortByPrice(flights []model.Flight) []model.Flight {
	sort.Slice(flights, func(i, j int) bool {
		return flights[i].Price < flights[j].Price
	})
	return flights
}
