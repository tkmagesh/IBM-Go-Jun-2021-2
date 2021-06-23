package main

import (
	"collections-demo/models"
	"fmt"
)

func main() {
	products := models.Products{}

	stove := models.Product{102, "Stove", 5000, 5, "Utencil"}
	anotherStove := models.Product{102, "Stove", 5000, 5, "Utencil"}
	fmt.Println(stove == anotherStove) //?

	products.AddProduct(models.Product{105, "Pen", 5, 50, "Stationary"})
	products.AddProduct(models.Product{107, "Pencil", 2, 100, "Stationary"})
	products.AddProduct(models.Product{103, "Marker", 50, 20, "Utencil"})
	products.AddProduct(models.Product{102, "Stove", 5000, 5, "Utencil"})
	products.AddProduct(models.Product{101, "Kettle", 2500, 10, "Stationary"})
	products.AddProduct(models.Product{104, "Scribble Pad", 20, 20, "Stationary"})

	fmt.Println("Products List")
	fmt.Println(products)

	fmt.Println("Index of stove => ", products.IndexOf(stove))
	fmt.Println("Is stove included in the list => ", products.Includes(stove))

	stationaryProductPredicate := func(product models.Product) bool {
		return product.Category == "Stationary"
	}
	fmt.Println("Are there any stationary prducts ? => ", products.Any(stationaryProductPredicate))

	foodProductPredicate := func(product models.Product) bool {
		return product.Category == "Food"
	}
	fmt.Println("Are there any food prducts ? => ", products.Any(foodProductPredicate))

	fmt.Println("Are all stationary prducts ? => ", products.All(stationaryProductPredicate))
	fmt.Println("Are all costly prducts ? => ", products.All(func(product models.Product) bool {
		return product.Cost > 50
	}))
	fmt.Println("Are there any costly prducts ? => ", products.Any(func(product models.Product) bool {
		return product.Cost > 50
	}))

	stationaryProducts := products.Filter(stationaryProductPredicate)
	fmt.Println("All stationary products from the list")
	fmt.Println(stationaryProducts)

	fmt.Println("Default sorting")
	products.Sort()
	fmt.Println(products)

	fmt.Println("Sorting by name")
	products.SortByName()
	fmt.Println(products)

	fmt.Println("Sorting by cost")
	products.SortByCost()
	fmt.Println(products)

}
