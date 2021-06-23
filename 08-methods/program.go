package main

import (
	"fmt"
	"strconv"
)

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

//using composition
/* type SpecialProduct struct {
	Product
} */

//the above using type aliasing
type SpecialProduct Product

//type aliasing
type MyStr string

//method for MyStr
func (myStr MyStr) WhoAmI() string {
	return fmt.Sprintf("I am string (%s)", myStr)
}

func main() {
	//var p Product
	//var p Product = Product{}
	//p := Product{}
	//p := Product{Id: 102, Name: "Pen", Cost: 5, Units: 50, Category: "Stationary"}
	p := Product{102, "Pen", 5, 50, "Stationary"}
	fmt.Println(p)
	//using the utility function
	//applyDiscount(&p, 10)

	//the above achieved using the method
	//(&p).applyDiscount(10)
	//A method can be invoked using the value OR pointer reference
	p.applyDiscount(10)
	fmt.Println("After applying 10% discount, product => ", p)

	//struct composition
	//grapes := PerishableProduct{Prod : Product{105, "Grapes", 50, 20, "Food"}, Expiry: "2 Days"}
	grapes := PerishableProduct{Product{105, "Grapes", 50, 20, "Food"}, "2 Days"}
	fmt.Println("Perishable Product (grapes) => ", grapes)

	//applying discount for perishable product
	//using the utility function
	//applyDiscount(&(grapes.Product), 10)
	//the above achieved using the method
	grapes.applyDiscount(10)
	fmt.Println("After applying 10% discount for Perishable product (grapes) => ", grapes)

	//using struct composition
	//specialPen := SpecialProduct{p}

	//using type aliasing
	specialPen := SpecialProduct(p)
	fmt.Println("Price of special pen => ", specialPen.GetPrice())
	//Note : In type aliasing, methods of the base type will NOT be available for the aliased type

	//type aliasing
	s := MyStr("Some dummy string")
	fmt.Println(s.WhoAmI())

	noAsString := MyStr("100")
	fmt.Println(strconv.Atoi(string(noAsString)))
}

//utility function
func applyDiscount(product *Product, discount float32) {
	product.Cost = product.Cost * ((100 - discount) / 100)
}

//utility function converted into a method
func (product *Product) applyDiscount(discount float32) {
	product.Cost = product.Cost * ((100 - discount) / 100)
}

func (sp SpecialProduct) GetPrice() float32 {
	return sp.Cost * 0.5
}
