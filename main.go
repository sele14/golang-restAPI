package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GOAL:
// an API that stores your investment portfolio.
// Allows for post, get, and delete requests of assets.

// (Note: This is file is more highly commented than usual for my own benefit.)
// ___________________________

// Create the different fields we expect for our investments (struct)
// The struct is our 'model' of our data (similar to a class)
type Instrument struct {
	ID       string  `json:"id"`
	Type     string  `json:"type"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
}

// declare the instruments var
var instruments []Instrument

func root(w http.ResponseWriter, r *http.Request) {
	// Fprintf prints the string to the writer object, not to console (it will show in the browser)
	fmt.Fprintf(w, "Hello, you are at: root()")
}

func getInstruments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Println prints to the console, not browser
	fmt.Println("Function Called: getInstruments()")

	// list instruments to the write stream (in the browser)
	json.NewEncoder(w).Encode(instruments)
}

func addInstruments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var instrument Instrument

	// convert inputted data to object we can work with
	json.NewDecoder(r.Body).Decode(&instrument)

	// Store the data we add
	instruments = append(instruments, instrument)
	// Encode the data and return it back, so user can verify stored data
	json.NewEncoder(w).Encode(instrument)
}

func deleteInstruments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get params from read stream
	params := mux.Vars(r)

	// pass the ID into the delete function
	deleteByID(params["id"])

	json.NewEncoder(w).Encode(instruments)
}

func deleteByID(id string) {
	// iterate by instruments list
	for index, instrument := range instruments {
		// if ID matches; delete
		if instrument.ID == id {
			instruments = append(instruments[:index], instruments[index+1:]...)
			break
		}

	}
}

func makeRequests() {
	// create router with short var (auto type, dont have to declare type)
	router := mux.NewRouter().StrictSlash(true)
	// Create a GET request
	router.HandleFunc("/", root).Methods("GET")

	// create get/post/delete requests to interact with the API
	router.HandleFunc("/instruments", getInstruments).Methods("GET")
	router.HandleFunc("/instruments", addInstruments).Methods("POST")
	router.HandleFunc("/instruments/{id}", deleteInstruments).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

// Some example asset classes:
// Equities
// Fixed Income
// Cash & Equiv
// Real Estate
// Commodities
// Derivatives
// Cryptocurrencies

func main() {
	// add our first instrument
	instruments = append(instruments, Instrument{
		ID:       "0",
		Type:     "Stock",
		Name:     "ATVI",
		Price:    92.68,
		Quantity: 5,
	})
	// add another instrument
	instruments = append(instruments, Instrument{
		ID:       "1",
		Type:     "Bond",
		Name:     "TNX",
		Price:    11.39,
		Quantity: 3,
	})
	makeRequests()
}
