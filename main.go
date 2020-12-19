package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// GOAL:
// an API that stores the ingredients for beer--e.g. an inventory for a beer factory.
// Allows for post, get, and delete requests of beer ingredients.

// (Note: This is file is more highly commented than usual for my own benefit.)

// Create the different fields we expect for our beer ingredients (struct)
// The struct is our 'model' of our data (similar to a class)

type Ingredient struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Quantity int `"json:"quantity"`
}

// declare the ingredients var

var ingredients []Ingredient

func root(w http.ResponseWriter, r *http.Request){
	// Fprintf prints the string to the writer object, not to console (it will show in the browser)
	fmt.Fprintf(w, "Hello, you are at: root()")
}

func getIngredients(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	// Println prints to the console, not browser
	fmt.Println("Function Called: getIngredients()")

	// list ingredients to the write stream (in the browser)
	json.NewEncoder(w).Encode(ingredients)
}

func addIngredients(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var ingredient Ingredient
	// Note to self: The _ used as var name tells the compiler to effectively discard
	// the RHS value, but to type-check it and evaluate it if it has any side effects.
	_ = json.NewDecoder(r.Body).Decode(&ingredient)
	// add to ingredients list
	
	// Store the data we add
	ingredients = append(ingredients, ingredient)
	// Encode the data and return it back, so user can verify stored data
	json.NewEncoder(w).Encode(ingredient)
}

func deleteIngredients(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	// Get params from read stream
	params := mux.Vars(r)

	// pass the ID into the delete function
	deleteByID(params["id"])

	json.NewEncoder(w).Encode(ingredients)
}

func deleteByID(id string){
	// iterate by ingredient list
	for index, ingredient := range ingredients {
		// if ID matches; delete
		if ingredient.ID == id {
			ingredients = append(ingredients[:index], ingredients[index+1:]...)
			break
		}

	}
}

func makeRequests(){
	// create router with short var (auto type, dont have to declare type) 
	router := mux.NewRouter().StrictSlash(true)
	// Create a GET request
	router.HandleFunc("/", root).Methods("GET")

	// create get/post/delete requests to interact with the API
	router.HandleFunc("/ingredients", getIngredients).Methods("GET")
	router.HandleFunc("/ingredients", addIngredients).Methods("POST")
	router.HandleFunc("/ingredients/{id}", deleteIngredients).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main(){
	// add water as ingredient
	ingredients = append(ingredients, Ingredient{
		ID : "0",
		Name : "Water",
		Desc : "Water is the largest volume ingredient in beer, and has a significant impact on the end product.",
		Quantity : 50,
	})
	// add another ingredient
	ingredients = append(ingredients, Ingredient{
		ID : "1",
		Name : "Hops",
		Desc : "Hops are responsible for producing aromas, some flavors and bitterness.",
		Quantity : 20,
	})
	// some more sample ingredients--can be added with Postman calls
	// ingredients = append(ingredients, Ingredient{
	// 	ID : "2",
	// 	Name : "Barley",
	// 	Desc : "Many brewers see barley as beer's soul. Barley has distinctive characteristics that make it a favored ingredient over other grains.",
	// 	Quantity : 20,
	// })
	// // add another ingredient
	// ingredients = append(ingredients, Ingredient{
	// 	ID : "3",
	// 	Name : "Yeast",
	// 	Desc : "Before it begins to reproduce and provide the beer with alcohol, yeast requires sugar to digest and oxygen to breathe.",
	// 	Quantity : 10,
	// })
	makeRequests()
}
