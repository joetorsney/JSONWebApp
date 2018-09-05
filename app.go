package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Panics if there is an error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Filepath should be relative to app.go
var products = initApp("prices.csv")

func searchHandler(w http.ResponseWriter, r *http.Request) {

	// If a panic is raised, respond with status 500.
	defer func() {
		if r := recover(); r != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	// Get the url query and call search products
	query := r.URL.Query()
	matches := searchProducts(query["name"][0], products)

	// Create json and check for errors
	js, err := json.Marshal(matches)
	check(err)

	// Finally, respond with the json
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/search", searchHandler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
