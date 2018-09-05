package main

// Definition for the product struct
// JSON can only marshal exported fields.
type Product struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}
