package repo

type FlightRepository interface {
	GetFlights() ([]Flight, error)
}
