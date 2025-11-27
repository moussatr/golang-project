package service

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"main.go/model"
	"main.go/repo"
)

type mockRepo struct {
	flights []model.Flight
	err     error
}

func (m *mockRepo) GetFlights() ([]model.Flight, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.flights, nil
}

func TestGetAllFlightsSortByPrice(t *testing.T) {
	service := &FlightService{
		Repos: []repo.FlightRepository{
			&mockRepo{flights: []model.Flight{
				{BokingId: "B1", Price: 350},
				{BokingId: "B2", Price: 150},
			}},
			&mockRepo{flights: []model.Flight{
				{BokingId: "B3", Price: 200},
			}},
		},
	}

	flights, err := service.GetAllFlights("price")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertBookingOrder(t, flights, []string{"B2", "B3", "B1"})
}

func TestGetAllFlightsSortByDeparture(t *testing.T) {
	departTimes := []time.Time{
		time.Date(2024, time.January, 1, 10, 0, 0, 0, time.UTC),
		time.Date(2024, time.January, 1, 7, 0, 0, 0, time.UTC),
		time.Date(2024, time.January, 1, 9, 30, 0, 0, time.UTC),
	}

	service := &FlightService{
		Repos: []repo.FlightRepository{
			&mockRepo{flights: []model.Flight{
				{BokingId: "B1", DepartureTime: departTimes[0]},
				{BokingId: "B2", DepartureTime: departTimes[1]},
			}},
			&mockRepo{flights: []model.Flight{
				{BokingId: "B3", DepartureTime: departTimes[2]},
			}},
		},
	}

	flights, err := service.GetAllFlights("departure")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertBookingOrder(t, flights, []string{"B2", "B3", "B1"})
}

func TestGetAllFlightsDefaultSort(t *testing.T) {
	service := &FlightService{
		Repos: []repo.FlightRepository{
			&mockRepo{flights: []model.Flight{
				{BokingId: "B1", Price: 500},
				{BokingId: "B2", Price: 100},
			}},
		},
	}

	flights, err := service.GetAllFlights("unknown")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertBookingOrder(t, flights, []string{"B2", "B1"})
}

func TestGetAllFlightsRepoError(t *testing.T) {
	service := &FlightService{
		Repos: []repo.FlightRepository{
			&mockRepo{err: errors.New("boom")},
		},
	}

	flights, err := service.GetAllFlights("price")

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if flights != nil {
		t.Fatalf("expected nil flights, got %v", flights)
	}
}

func assertBookingOrder(t *testing.T, flights []model.Flight, expected []string) {
	t.Helper()

	order := make([]string, len(flights))
	for i, f := range flights {
		order[i] = f.BokingId
	}

	if !reflect.DeepEqual(order, expected) {
		t.Fatalf("expected booking order %v, got %v", expected, order)
	}
}

