package repo

import "time"

type Flight struct {
	BokingId         string    `json:"bookingId"`
	Reference        string    `json:"reference"`
	Traveler         traveler  `json:"traveler"`
	Total            total     `json:"total"`
	Segments         []segment `json:"segments"`
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

type traveler struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type total struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type segment struct {
	Flight flight `json:"flight"`
}

type flight struct {
	Number string    `json:"number"`
	From   string    `json:"from"`
	To     string    `json:"to"`
	Depart time.Time `json:"depart"`
	Arrive time.Time `json:"arrive"`
}
