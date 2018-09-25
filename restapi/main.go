package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"log"
)

type Car struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Wheel string `json:"wheel"`
}
var car = Car{"benz", "SUV", "4"}

func ServerCar(writer http.ResponseWriter, request *http.Request)  {
	json.NewEncoder(writer).Encode(car)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/car", ServerCar).Methods("GET")
	log.Fatal(http.ListenAndServe(":8999", router))
}
