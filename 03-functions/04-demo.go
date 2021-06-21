//higher order functions
package main

import "fmt"

func main() {

	//functions as arguments and return values
	loggerAdd := log(add)
	loggerSubtract := log(subtract)
	loggerMultiply := log(multiply)
	loggerDivide := log(divide)
	fmt.Println(loggerAdd(100, 200))
	fmt.Println(loggerSubtract(100, 200))
	fmt.Println(loggerMultiply(100, 200))
	fmt.Println(loggerDivide(100, 200))

}

func log(fn func(int, int) int) func(int, int) int {
	return func(x, y int) int {
		fmt.Printf("Operation performed for %d and %d\n", x, y)
		result := fn(x, y)
		return result
	}
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
