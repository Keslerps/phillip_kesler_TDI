package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var myClient = &http.Client{Timeout: 2 * time.Second}

func handler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	fmt.Fprintf(w, "Search term: %s\n\n", params["search"])
	req, err := http.NewRequest("GET", "http://www.reddit.com/r/anime.json", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	req.Header.Set("User-Agent", "TDI_Application")
	re, err := myClient.Do(req)
	if err != nil {
		println("Come ON")
	}
	target := new(redditAPI)
	body, erro := ioutil.ReadAll(re.Body)
	if erro != nil {
		panic(erro.Error())
	}
	unErr := json.Unmarshal(body, target)
	if unErr != nil {
		panic(unErr.Error())
	}
	defer re.Body.Close()
	output(w, target.Data, params["search"])

}
