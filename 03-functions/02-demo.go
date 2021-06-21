package main

import "fmt"

func main() {
	fmt.Println(sum(10, 20))
	fmt.Println(sum(10, 20, 30, 40))
	fmt.Println(sum())
	fmt.Println(sum(10))
}

func sum(nos ...int) int {
	result := 0
	for idx := 0; idx < len(nos); idx++ {
		result += nos[idx]
	}
	return result
}
