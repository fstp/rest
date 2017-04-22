package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Person doc
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address doc
type Address struct {
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
}

var people []Person

func getPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func getPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func createPersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func deletePersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func main() {
	people = append(people, Person{
		ID:        "1",
		Firstname: "John",
		Lastname:  "Doe",
		Address: &Address{
			City:    "Stockholm",
			Country: "Sweden"}})

	people = append(people, Person{
		ID:        "2",
		Firstname: "Jane",
		Lastname:  "Doe"})

	router := mux.NewRouter()
	router.HandleFunc("/people", getPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", getPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", createPersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", deletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))
}
