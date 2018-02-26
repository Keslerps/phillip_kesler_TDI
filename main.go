package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/redditer/:{search:[a-zA-Z]+}", handler).Methods("GET")
	http.Handle("/", rtr)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
