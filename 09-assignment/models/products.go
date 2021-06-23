package models

import "fmt"

type Products []Product

//implement sorting for products (should use sort.Sort())
//Default sort - should sort the products by id
//Sort by name
//Sort by cost
//Sort by Units
//Sort by Category

func (products *Products) AddProduct(product Product) {
	*products = append(*products, product)
}

func (products Products) IndexOf(product Product) int {
	for idx, prod := range products {
		if prod == product {
			return idx
		}
	}
	return -1
}

func (products Products) Includes(product Product) bool {
	return products.IndexOf(product) != -1
}

func (products Products) Any(predicate func(Product) bool) bool {
	for _, prod := range products {
		if predicate(prod) {
			return true
		}
	}
	return false
}

func (products Products) All(predicate func(Product) bool) bool {
	for _, prod := range products {
		if !predicate(prod) {
			return false
		}
	}
	return true
}

func (products Products) String() string {
	result := ""
	for _, product := range products {
		result += fmt.Sprint(product)
	}
	return result
}

func (products Products) Filter(predicate func(Product) bool) *Products {
	result := &Products{}
	for _, prod := range products {
		if predicate(prod) {
			*result = append(*result, prod)
		}
	}
	return result
}
