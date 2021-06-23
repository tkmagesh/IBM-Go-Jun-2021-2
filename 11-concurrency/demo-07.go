package main

import (
	"fmt"
	"time"
)

func main() {
	print("Hello", 2)
	print("World", 4)
	var input string
	fmt.Scanln(&input)
}

func print(msg string, delay time.Duration /* .... */) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(delay * time.Second)
	}
}

/*
Hello
World
Hello
World
Hello
World
Hello
World
Hello
World
*/
