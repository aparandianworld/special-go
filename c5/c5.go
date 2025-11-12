package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("module: c5")
	fmt.Println("version: 0.0.1")

	var result strings.Builder
	jsonData := `[
    {"id": 1, "name": "Laptop", "price": 999.99, "category": "Electronics"},
    {"id": 2, "name": "Headphones", "price": 199.99, "category": "Electronics"},
    {"id": 3, "name": "Coffee Maker", "price": 79.99, "category": "Appliances"},
    {"id": 4, "name": "Gaming Console", "price": 499.99, "category": "Electronics"}
]`
	products, err := parseProductCatalog(jsonData)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		for _, p := range products {
			result.WriteString(fmt.Sprintf("%v\n", p))
		}
	}

	fmt.Println("products: \n", result.String())

}

type Product struct {
	ID       int
	Name     string
	Price    float64
	Category string
}

func parseProductCatalog(jsonData string) ([]Product, error) {

	var products []Product

	err := json.Unmarshal([]byte(jsonData), &products)

	if err != nil {
		return nil, err
	}

	return products, nil
}
