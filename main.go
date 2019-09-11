package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Pet Structure
type Pet struct {
	ID          string `json:"Id"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Kind        string `json:"Kind"`
	Age         int    `json:"Age"`
	Attributes  string `json:"Attributes"`
}

// Pets to Create a new type
// which is a slice of strings
var Pets []Pet

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is Running!")
	fmt.Println("Health Enpoint: health")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Homepage")
	json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to AllPets"})
}

func petsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: petsHandler")
	json.NewEncoder(w).Encode(Pets)
}

func petsDetailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: petsDetailHandler")
	vars := mux.Vars(r)
	param := vars["id"]

	fmt.Print("Param: " + param)

	for _, pet := range Pets {
		if pet.ID == param {
			json.NewEncoder(w).Encode(pet)
		}
	}
}

func petsCreateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: petsCreateHandler")
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(reqBody))
	var pet Pet
	json.Unmarshal(reqBody, &pet)
	Pets = append(Pets, pet)

	json.NewEncoder(w).Encode(pet)
}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/health", health)
	p := router.PathPrefix("/pets").Subrouter()
	p.HandleFunc("", petsCreateHandler).Methods("POST")
	p.HandleFunc("", petsHandler)
	p.HandleFunc("/{id}/details", petsDetailHandler)
	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	Pets = []Pet{
		Pet{ID: "1", Name: "Pixie", Description: "A small furry friend", Kind: "Dog", Age: 7, Attributes: "Brown, Sweet, Tiny"},
		Pet{ID: "2", Name: "Sombra", Description: "A wild explorer", Kind: "Cat", Age: 2, Attributes: "Black, Playful, Wild"},
	}
	handleRequest()
}
