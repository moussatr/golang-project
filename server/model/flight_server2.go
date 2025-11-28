package model

import "time"

// FlightServer2 represents the raw data structure from j-server2
type FlightServer2 struct {
	Reference string           `json:"reference"`
	Status    string           `json:"status"`
	Traveler  TravelerServer2  `json:"traveler"`
	Segments  []SegmentServer2 `json:"segments"`
	Total     TotalServer2     `json:"total"`
}

type TravelerServer2 struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type TotalServer2 struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type SegmentServer2 struct {
	Flight FlightSegmentServer2 `json:"flight"`
}

type FlightSegmentServer2 struct {
	Number string    `json:"number"`
	From   string    `json:"from"`
	To     string    `json:"to"`
	Depart time.Time `json:"depart"`
	Arrive time.Time `json:"arrive"`
}
