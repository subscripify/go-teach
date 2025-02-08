package main

import (
	"encoding/json"
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
)

type Message struct {
	Message string `json:"message"`
}

type ResponseMessage struct {
	Message     string  `json:"message"`
	YourMessage Message `json:"your_message"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/message", messageHandler).Methods("GET")
	r.HandleFunc("/echo", echoHandler).Methods("POST")
	log.Println("Server starting on port 8080...")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("could not start listener %v", err.Error())
	}

}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	response := Message{Message: "Hello World"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	var msg Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	response := ResponseMessage{
		Message:     "I processed your message",
		YourMessage: msg,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
