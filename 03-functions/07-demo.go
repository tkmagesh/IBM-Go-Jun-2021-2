//panic & recovery
package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("Exiting from main")
		r := recover()
		if r != nil {
			//app is recovering from the panic
			fmt.Println("Something went wrong! Contact the administrator..!!")
			fmt.Println(r)
			//debug.PrintStack()
		}
	}()
	//fmt.Println(add(100, 200))
	quotient, remainder := divideClient(100, 0)
	fmt.Println(quotient, remainder)

}

func add(x, y int) int {
	return x + y
}

func divideClient(x, y int) (quotient, remainder int) {
	defer func() {
		panic("panic from the deferred fn of divideClient")
	}()
	return divide(x, y)
}

func divide(x, y int) (quotient, remainder int) {
	/* if y == 0 {
		panic("Attempt to divide by zero!")
	} */
	quotient = x / y
	remainder = x % y
	return
}
