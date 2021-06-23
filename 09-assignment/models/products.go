package models

import (
	"fmt"
	"sort"
)

type Products []Product

//implement sorting for products (should use sort.Sort())
//Default sort - should sort the products by id
//Sort by name
//Sort by cost
//Sort by Units
//Sort by Category

func (products Products) Len() int {
	return len(products)
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

type SortByName struct {
	Products
}

func (sortByName SortByName) Less(i, j int) bool {
	return sortByName.Products[i].Name < sortByName.Products[j].Name
}

type SortByCost struct {
	Products
}

func (sortByCost SortByCost) Less(i, j int) bool {
	return sortByCost.Products[i].Cost < sortByCost.Products[j].Cost
}

// Swap swaps the elements with indexes i and j.

func (products Products) Sort() {
	sort.Sort(products)
}

func (products Products) SortByName() {
	sort.Sort(SortByName{products})
}

func (products Products) SortByCost() {
	sort.Sort(SortByCost{products})
}

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
