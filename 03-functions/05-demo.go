//closures
package main

import "fmt"

func main() {
	//counter := getCounter()
	counter := func() func() int {
		var count int
		return func() int {
			count += 1
			return 1
		}
	}()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func getCounter() func() int {
	var count int
	return func() int {
		count += 1
		return count
	}
}
