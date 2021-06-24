package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fibonacci(ch)
	go printFibonacci(ch)
	fmt.Println("Press ENTER to stop!")
	var input string
	fmt.Scanln(&input)
	fmt.Println("Thats all folks!")
}

func printFibonacci(ch chan int) {
	for no := range ch {
		fmt.Println(no)
	}
}

func fibonacci(ch chan int) {
	x, y := 0, 1
	stop := time.After(10 * time.Second)
	for {
		select {
		case ch <- x:
			time.Sleep(1 * time.Second)
			x, y = y, x+y
		case <-stop:
			fmt.Println("Quitting")
			close(ch)
			return
		}
	}
}
