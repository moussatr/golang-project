package model

import "time"

// Flight represents the normalized structure used by the aggregator.
// It matches the shape returned by j-server1 so downstream code can treat
// flights from both sources identically.
type Flight struct {
	BokingId         string    `json:"bookingId"`
	Status           string    `json:"status"`
	PassengerName    string    `json:"passengerName"`
	FlightNumber     string    `json:"flightNumber"`
	DepartureAirport string    `json:"departureAirport"`
	ArrivalAirport   string    `json:"arrivalAirport"`
	DepartureTime    time.Time `json:"departureTime"`
	ArrivalTime      time.Time `json:"arrivalTime"`
	Price            float64   `json:"price"`
	Currency         string    `json:"currency"`
}
