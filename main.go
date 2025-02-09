package main

import (
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

}

func colorMaker(w http.ResponseWriter, r *http.Request) {

}
