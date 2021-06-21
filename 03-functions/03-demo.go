//higher order functions
package main

import "fmt"

func main() {

	/* fn := func() {
		fmt.Println("fn is invoked")
	} */
	var fn func() = func() {
		fmt.Println("fn is invoked")
	}
	fn()
	//anonymous function
	func(x int) {
		fmt.Println("anonymous function invoked with arg : ", x)
	}(100)

	/*
		var compute func(int, int) int
		compute = func(x, y int) int {
			return x + y
		}
		fmt.Println(compute(100, 200))

		compute = func(x, y int) int {
			return x - y
		}
		fmt.Println(compute(100, 50))
	*/

	//passing functions as argument to another function
	fmt.Println(compute(add, 100, 200))
	fmt.Println(compute(subtract, 100, 200))
	fmt.Println(compute(multiply, 100, 200))
	fmt.Println(compute(divide, 100, 200))
}

func compute(fn func(int, int) int, x int, y int) int {
	fmt.Printf("Operation performed for %d and %d\n", x, y)
	result := fn(x, y)
	return result
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}
