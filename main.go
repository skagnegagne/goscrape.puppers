/*
Web scraper, scraping dog facts (not all have to do with dogs; very disapointed)
API to find and filter facts from site.
*/
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//defining fact structure
type Fact struct {
	ID   int    `json:"id"`
	Fact string `json:"fact"`
}

//setting up routes and starting API server
func main() {

	Router := MountRoutes()
	if err := http.ListenAndServe(":3000", Router); err != nil {
		log.Fatal(err)
	}

}

//converting fact slice to JSON data
func writeJSON(data []Fact) string {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return ""
	}
	return string(file)
}
