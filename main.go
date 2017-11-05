package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//The person type
type Person struct {
	ID        string   `json:"id.omitempty"`
	Firstname string   `json:"firstname.omitempty"`
	Lastname  string   `json:"lastname.omitempty"`
	Address   *Address `json:"address.omitempty"`
}

type Address struct {
	City  string `json:"city.omitempty"`
	State string `json:"state.omitempty"`
}

var people []Person

//Display aall from the people var

//make function with http's method funcName(write with http.Response, read with *http.Request)

//Display all from the people var
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

//Display single user data
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
		json.NewEncoder(w).Encode(&Person{})
	}
}

//create new item
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

//Delete an specific item
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
}

//main function for routing for request and response
func main() {
	//we will do response for request then we need to set route rule
	//first assign route variable with mux's method
	router := mux.NewRouter()
	//create data with manual
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Alan", Lastname: "P", Address: &Address{City: "City Y", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Kartick", Lastname: "karan", Address: &Address{City: "City Z", State: "State Z"}})
	//we make API endpoint like below
	/*
		/people(GET) -> all persons in the phonebook document(database)
		/poople/{id}(GET) -> Get a single person for id
		/people/{id}(POST) -> Create a new persons for id
		/people/{id}(DELETE) -> Delete a person for id
	*/
	//we make route handle for below endpoint
	//route.HandleFunc has two parameter one is endpoint(routing target) the other is function
	//then with method type
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	//then we will make server for listining of request with log.Fatal
	//log.Fatal has two parameter, one is portnumber theother is router variable
	log.Fatal(http.ListenAndServe(":8000", router))
}
