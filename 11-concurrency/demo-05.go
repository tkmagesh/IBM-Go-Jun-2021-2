package main

import (
	"fmt"
)

//"share memory by communicating (by using channels)"

func main() {
	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch //read the data from the channel
	fmt.Println(result)
}

func add(x, y int, ch chan int) {
	result := x + y
	ch <- result //write data into the channel
}
