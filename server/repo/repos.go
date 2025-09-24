package repo

type FlightRepository interface {
	GetFlightsByPrice(minPrice, maxPrice float64) ([]Flight, error)
	GetAllFlights() ([]Flight, error)
}
