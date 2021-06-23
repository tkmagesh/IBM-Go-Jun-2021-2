package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go print("Hello", 1, ch1, ch2)
	go print("World", 1, ch2, ch1)
	ch1 <- "start"
	var input string
	fmt.Scanln(&input)
}

func print(msg string, delay time.Duration, in chan string, out chan string) {
	for i := 0; i < 5; i++ {
		<-in
		fmt.Println(msg)
		time.Sleep(delay * time.Second)
		out <- "done"
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
