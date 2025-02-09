package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	mux "github.com/gorilla/mux"
)

type RequestMessage struct {
	Color string `json:"color"`
}

type ResponseMessage struct {
	Color   string `json:"color"`
	Message string `json:"message"`
}

// this is your data store - in the real world it would be a database
var (
	dataStore []string
	mu        sync.Mutex
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/color/{color}", colorRetriever).Methods("GET")
	r.HandleFunc("/color", colorMaker).Methods("POST")
	log.Println("Server starting on port 8080...")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("could not start listener %v", err.Error())
	}
}

func colorRetriever(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request to get color received")
	color, err := mux.Vars(r)["color"]

	if !err {
		response := ResponseMessage{
			Message: "Invalid request",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	fmt.Printf("getting %s \n", color)
	for _, c := range dataStore {
		if c == color {
			response := ResponseMessage{
				Message: color + " is a nice color",
				Color:   color,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	response := ResponseMessage{
		Message: color + " is not found",
		Color:   color,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)
}

func colorMaker(w http.ResponseWriter, r *http.Request) {

}
