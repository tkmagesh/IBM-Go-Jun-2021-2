package main

import "fmt"

func main() {
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func counter() int {
	var count int
	count += 1
	return count
}
