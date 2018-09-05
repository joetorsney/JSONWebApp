package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

// Loads the CSV file and returns a slice of product structs
func initApp(filepath string) []Product {

	// Create product struct
	products := []Product{}

	// Open file
	f, err := os.Open(filepath)
	defer f.Close()
	check(err)

	reader := csv.NewReader(bufio.NewReader(f))
	// Start reading
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else {
			check(err)
		}

		products = append(products, Product{
			Name:  line[0],
			Price: line[1],
		})
	}

	return products
}
