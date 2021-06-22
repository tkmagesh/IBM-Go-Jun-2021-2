//error handling
package main

import (
	"errors"
	"fmt"
)

func main() {

	quotient, remainder, err := divide(100, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(quotient, remainder)

}

func add(x, y int) int {
	return x + y
}

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = errors.New("Attempt to divide by zero")
		return
	}
	quotient = x / y
	remainder = x % y
	return
}
