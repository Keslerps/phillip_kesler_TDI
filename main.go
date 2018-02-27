package main

import (
	"log"
	"net/http"
)

/**
 * Main entry point for the API handler. Calls
 * startUp() to start the router on port 8080 for
 * all incoming requests.
 **/
func main() {
	log.Fatal(http.ListenAndServe(":8080", startUp()))
}
