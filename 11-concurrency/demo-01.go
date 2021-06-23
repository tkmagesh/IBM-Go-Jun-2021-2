package main

import (
	"fmt"
	"time"
)

func main() {
	go print("Hello")
	print("World")

	/*
		var input string
		fmt.Scanln(&input)
	*/
	time.Sleep(1 * time.Millisecond)
}

func print(msg string) {
	fmt.Println(msg)
}
