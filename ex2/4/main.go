package main

import (
	"encoding/json"
	"fmt"
)

// Define the Product struct
type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

// Function to encode a Product to JSON
func EncodeProduct(p Product) (string, error) {
	jsonData, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// Function to decode JSON into a Product
func DecodeProduct(data string) (Product, error) {
	var p Product
	err := json.Unmarshal([]byte(data), &p)
	return p, err
}

func main() {
	product := Product{Name: "Laptop", Price: 999.99, Quantity: 5}

	jsonStr, err := EncodeProduct(product)
	if err != nil {
		fmt.Println("Error encoding:", err)
		return
	}
	fmt.Println("JSON:", jsonStr)

	decodedProduct, err := DecodeProduct(jsonStr)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}
	fmt.Printf("Decoded Product: %+v\n", decodedProduct)
}
