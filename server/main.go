package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
	"main.go/controller"
	"main.go/repo"
	"main.go/service"
)

func main() {
	fmt.Println("starting point !")

	viper.AutomaticEnv()

	repo1 := &repo.Server1Repo{BaseURL: viper.GetString("BASE_URL1") + ":" + viper.GetString("JSERVER1_PORT")}
	repo2 := &repo.Server2Repo{BaseURL: viper.GetString("BASE_URL2") + ":" + viper.GetString("JSERVER2_PORT")}

	fmt.Println("Repository 1 URL:", repo1.BaseURL)
	fmt.Println("Repository 2 URL:", repo2.BaseURL)
	fmt.Println("Server 1 Port:", viper.GetString("BASE_URL"))

	flightService := &service.FlightService{Repos: []repo.FlightRepository{repo1, repo2}}
	controller := &controller.FlightController{Service: flightService}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", controller.Health)
	mux.HandleFunc("/flights", controller.GetFlights)

	http.ListenAndServe(":3001", mux)

}
