package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/**
 * Main entry point for the API handler. Creates
 * a new Mux Router (from github.com/gorilla/mux)
 * and handles any request in the form of localhost:8000/redditer/:
 * followed by a search term, any case.
 **/
func main() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/redditer/:{search:[a-zA-Z]+}", handler).Methods("GET")
	http.Handle("/", rtr)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
