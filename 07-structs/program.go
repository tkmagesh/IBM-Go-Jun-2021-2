package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

//struct composition
//part - 1
/*
type PerishableProduct struct {
	Prod Product
	Expiry  string
}
*/

type PerishableProduct struct {
	Product
	Expiry string
}

func main() {
	//var p Product
	//var p Product = Product{}
	//p := Product{}
	//p := Product{Id: 102, Name: "Pen", Cost: 5, Units: 50, Category: "Stationary"}
	p := Product{102, "Pen", 5, 50, "Stationary"}
	fmt.Println(p)
	applyDiscount(&p, 10)
	fmt.Println("After applying 10% discount, product => ", p)

	//struct composition
	//grapes := PerishableProduct{Prod : Product{105, "Grapes", 50, 20, "Food"}, Expiry: "2 Days"}
	grapes := PerishableProduct{Product{105, "Grapes", 50, 20, "Food"}, "2 Days"}
	fmt.Println("Perishable Product (grapes) => ", grapes)

	//applying discount for perishable product
	applyDiscount(&(grapes.Product), 10)
	fmt.Println("After applying 10% discount for Perishable product (grapes) => ", grapes)
}

//utility function
func applyDiscount(product *Product, discount float32) {
	product.Cost = product.Cost * ((100 - discount) / 100)
}
