package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fibonacci(10, ch)
	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Thats all folks!")
}

func fibonacci(count int, ch chan int) {
	x, y := 0, 1
	for idx := 0; idx < count; idx++ {
		time.Sleep(1 * time.Second)
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
