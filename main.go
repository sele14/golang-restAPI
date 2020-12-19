package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

// GOAL:
// an API that stores the ingredients for beer--e.g. an inventory for a beer factory.


// Create the different fields we expect(items for our beer ingredients in this case)

type Item struct {
	UID string `json:"UID"`
	Name string `json:"Name"`
	Desc string `json:"Desc"`
	Quantity int `"json:"Quantity"`
}

// declare the ingredients var

var ingredients []Item

func homePage(w http.ResponseWriter, r *http.Request){
	// Fprintf prints the string to the writer object, not to console (it will show in the browser)
	fmt.Fprintf(w, "Endpoint called: homePage()")
}

func getIngredients(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content Type", "application/json")
	// Println prints to the console, not browser
	fmt.Println("Function Called: getIngredients()")

	// list ingredients to the write stream (in the browser)
	json.NewEncoder(w).Encode(ingredients)
}

func addIngredients(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content Type", "application/json")
	var item Item
	// Note to self: The _ used as var name tells the compiler to effectively discard
	// the RHS value, but to type-check it and evaluate it if it has any side effects.
	_ = json.NewDecoder(r.Body).Decode(&item)
	// add to ingredients list
	
	// Store the data we add
	ingredients = append(ingredients, item)
	// Encode the data and return it back, so user can verify stored data
	json.NewEncoder(w).Encode(item)
}

func makeRequests(){
	// create router with short var (auto type, dont have to declare type) 
	router := mux.NewRouter().StrictSlash(true)
	// Create a GET request
	router.HandleFunc("/", homePage).Methods("GET")
	// another route to look at the ingredients
	router.HandleFunc("/ingredients", getIngredients).Methods("GET")
	// post request to add ingredients
	router.HandleFunc("/ingredients", addIngredients).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main(){
	// add water as ingredient
	ingredients = append(ingredients, Item{
		UID : "0",
		Name : "Water",
		Desc : "Water is the largest volume ingredient in beer, and has a significant impact on the end product.",
		Quantity : 50,
	})
	// add another ingredient
	ingredients = append(ingredients, Item{
		UID : "1",
		Name : "Hops",
		Desc : "Hops are responsible for producing aromas, some flavors and bitterness.",
		Quantity : 20,
	})
	// some more sample ingredients;
	// ingredients = append(ingredients, Item{
	// 	UID : "2",
	// 	Name : "Barley",
	// 	Desc : "Many brewers see barley as beer's soul. Barley has distinctive characteristics that make it a favored ingredient over other grains.",
	// 	Quantity : 20,
	// })
	// // add another ingredient
	// ingredients = append(ingredients, Item{
	// 	UID : "3",
	// 	Name : "Yeast",
	// 	Desc : "Before it begins to reproduce and provide the beer with alcohol, yeast requires sugar to digest and oxygen to breathe.",
	// 	Quantity : 10,
	// })
	makeRequests()
}
