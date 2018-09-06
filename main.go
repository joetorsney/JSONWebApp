package main

import (
	"encoding/json"
	"fmt"
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

func reqHandler(w http.ResponseWriter, r *http.Request) {

	// If a panic is raised, respond with status 500.
	defer func() {
		if r := recover(); r != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	// r.URL.Query returns a map[string][]string
	query := r.URL.Query()
	var matches []Product

	fmt.Println(r.URL.Path)
	switch r.URL.Path {
	case "/search":
		matches = searchProducts(query["name"][0], products)
	case "/fuzzysearch":
		matches = fuzzySearch(query["name"][0], products)
	default:
		// If the url is unsupported, eg /hello, respond with 404 and return.
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Create json and check for errors
	js, err := json.Marshal(matches)
	check(err)

	// Finally, respond with the json
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func main() {
	http.HandleFunc("/", reqHandler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
