package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

/**
 * Custom HTTP client used for requests. Will Timeout
 * after 2 seconds if a request can not be received.
 **/
var myClient = &http.Client{Timeout: 2 * time.Second}

/**
 * Creates a Gorilla Mux Router (from the provided GitHub repository)
 * and sets it up to handle /redditer/:search, where search can be
 * any word, ignoring case. Returns the made router on completion.
 **/
func startUp() *mux.Router {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/redditer/:{search:[a-zA-Z_]+}", handler).Methods("GET")
	return rtr
}

/**
 * HTTP handler for a request from main. Will issue a GET
 * reuqest to reddit.com/r/subReddit.json, with a request with a USER_AGENT
 * header as to not raise Reddit's bot protections, and then
 * Unmarshal the data into the Reddit_API struct before sending
 * that struct to output for writing to the response. If
 * any error is ever encountered, the program exits.
 * @param w Http.ReponseWriter where output it printed.
 * @param r pointer to a http.Request that stores the actual
 * request and search parameter.
 **/
func handler(w http.ResponseWriter, r *http.Request) {
	//Search parameter defined by the :search field
	params := mux.Vars(r)
	/**
	 * Create the request, be a good internet denizen by
	 * setting the header as a User-Agent, and issue
	 * the GET request.
	 **/
	req, err := http.NewRequest("GET", "http://www.reddit.com/r/"+subReddit+".json", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	req.Header.Set("User-Agent", "TDI_Application")
	re, err := myClient.Do(req)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	/**
	 * Reads all the data in the get request into
	 * a byte array called body, and error check.
	 **/
	target := new(redditAPI)
	body, err := ioutil.ReadAll(re.Body)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	/**
	 * unMarshall all the byte data in body into a
	 * Reddit_API struct to fill the .json fields.
	 **/
	unErr := json.Unmarshal(body, target)
	if unErr != nil {
		log.Print(unErr)
		os.Exit(1)
	}
	//Defer body close and send data to output
	defer re.Body.Close()
	output(w, target.Data, params["search"])

}
