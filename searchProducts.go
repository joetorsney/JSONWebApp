package main

import "strings"

// Given a monogram, e.g. "hello", this will return a slice of products whose
// name contains the monogram.
func searchProducts(query string, products []Product) []Product {
	matches := []Product{}

	// Iterate through products
	for i := 0; i < len(products); i++ {
		// Seperate words into a slice of strings
		monograms := strings.Split(products[i].Name, " ")

		// Iterate through monograms that were in the product name and test for match with query
		for j := 0; j < len(monograms); j++ {

			if query == monograms[j] {
				// If match is found, add the product to the matches slice and stop searching.
				matches = append(matches, products[i])
				break
			}
		}
	}
	return matches
}
