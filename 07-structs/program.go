package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

//composition
//part-1
/*
type PerishableProduct struct {
	product Product
	Expiry  string
}
*/

//part-2
type PerishableProduct struct {
	Product
	Expiry string
}

func main() {
	//var p Product
	//var p Product = Product{Id: 102, Name: "Pen", Cost: 5, Units: 20, Category: "Stationary"}
	var p Product = Product{102, "Pen", 5, 20, "Stationary"}
	fmt.Println(p)
	//applyDiscount(p, 10)
	p.applyDiscount(10)
	fmt.Println(p) //=> p.Cost should reflect the discount

	//composition
	//part-1
	/* grapes := PerishableProduct{product: Product{102, "Grapes", 70, 10, "Food"}, Expiry: "2 Days"} */
	/*
		grapes := PerishableProduct{Product{102, "Grapes", 70, 10, "Food"}, "2 Days"}
		fmt.Println(grapes)
		fmt.Println("Cost of grapes => ", grapes.product.Cost)
	*/
	//Part-2
	//grapes := PerishableProduct{Product: Product{102, "Grapes", 70, 10, "Food"}, Expiry: "2 Days"}
	grapes := PerishableProduct{Product{102, "Grapes", 70, 10, "Food"}, "2 Days"}
	//grapes := NewPerishableProduct(102, "Grapes", 70, 10, "Food", "2 Days")
	fmt.Println(grapes)
	//fmt.Println("Cost of grapes => ", grapes.Product.Cost)
	fmt.Println("Cost of grapes => ", grapes.Cost)

	//applying discount using the 'applyDiscount' function
	//applyDiscount(&(grapes.Product), 20)

	//applying discount using the 'applyDiscount' method
	grapes.applyDiscount(20)
	fmt.Println("After applying 20% discount, grapes => ", grapes)

}

//function
func applyDiscount(p *Product, discount float32) {
	//(*p).Cost = (*p).Cost * ((100 - discount) / 100)
	p.Cost = p.Cost * ((100 - discount) / 100)
}

//method
func (p *Product) applyDiscount(discount float32) {
	//(*p).Cost = (*p).Cost * ((100 - discount) / 100)
	p.Cost = p.Cost * ((100 - discount) / 100)
}

func NewPerishableProduct(id int, name string, cost float32, units int, category string, expiry string) *PerishableProduct {
	return &PerishableProduct{Product{id, name, cost, units, category}, expiry}
}
