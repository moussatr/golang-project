package utils

import "main.go/model"


func TransformServer2ToFlight(server2Flight model.FlightServer2) model.Flight {
	if len(server2Flight.Segments) == 0 {
		return model.Flight{
			BokingId:      server2Flight.Reference,
			Status:        server2Flight.Status,
			PassengerName: combineName(server2Flight.Traveler.FirstName, server2Flight.Traveler.LastName),
			Price:         server2Flight.Total.Amount,
			Currency:      server2Flight.Total.Currency,
		}
	}

	firstSegment := server2Flight.Segments[0].Flight
	lastSegment := server2Flight.Segments[len(server2Flight.Segments)-1].Flight

	return model.Flight{
		BokingId:         server2Flight.Reference,
		Status:           server2Flight.Status,
		PassengerName:    combineName(server2Flight.Traveler.FirstName, server2Flight.Traveler.LastName),
		FlightNumber:     firstSegment.Number,
		DepartureAirport: firstSegment.From,
		ArrivalAirport:   lastSegment.To,
		DepartureTime:    firstSegment.Depart,
		ArrivalTime:      lastSegment.Arrive,
		Price:            server2Flight.Total.Amount,
		Currency:         server2Flight.Total.Currency,
	}
}

func combineName(first, last string) string {
	switch {
	case first == "" && last == "":
		return ""
	case first == "":
		return last
	case last == "":
		return first
	default:
		return first + " " + last
	}
}
