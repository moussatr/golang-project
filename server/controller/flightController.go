package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"main.go/service"
)

type FlightController struct {
	Service *service.FlightService
}

func (c *FlightController) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OK"))
}

func (c *FlightController) GetFlights(w http.ResponseWriter, r *http.Request) {
	sortBy := r.URL.Query().Get("sort")
	fmt.Println("Sort by:", sortBy)
	flights, err := c.Service.GetAllFlights(sortBy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flights)
}
