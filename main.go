package main

import (
	"net/http"
	"sync"
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

}

func colorRetriever(w http.ResponseWriter, r *http.Request) {

}

func colorMaker(w http.ResponseWriter, r *http.Request) {

}
