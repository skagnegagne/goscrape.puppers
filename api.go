package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"
)

//returning all facts with ability to filter
func FactsHandler(w http.ResponseWriter, r *http.Request) {
	Facts := ScrapeFacts()
	//search function, check if search query is specified
	if r.FormValue("search") == "" {
		jsonFile := writeJSON(Facts)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", jsonFile)
	} else {
		factList := make([]Fact, 0)
		searcher, err := regexp.Compile(fmt.Sprintf("(?i)%s", r.FormValue("search")))
		if err != nil {
			w.Write([]byte("Server error searching"))
		}
		for _, v := range Facts {
			if searcher.MatchString(v.Fact) {
				factList = append(factList, v)
			}
		}
		jsonFile := writeJSON(factList)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", jsonFile)
	}
}

//returns single fact, chosen by fact ID
func FactHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Facts := ScrapeFacts()
	var factMatch Fact
	var foundMatch = false
	for _, v := range Facts {
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			continue
		}
		if v.ID == ID {
			factMatch = v
			foundMatch = true
		}
	}
	if foundMatch {
		result := []Fact{factMatch}

		jsonFile := writeJSON(result)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", jsonFile)
	} else {
		fmt.Fprintf(w, "Fact ID not found")
	}

}

//creates and configures router
func MountRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/facts", FactsHandler).Methods("GET")
	r.HandleFunc("/facts", FactsHandler).Methods("GET").Queries("search", "{search}")
	r.HandleFunc("/facts/{id}", FactHandler).Methods("GET")
	return r
}
