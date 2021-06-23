package main

import (
	"fmt"
	"sync"
)

//"share memory by communicating (by using channels)"

var wg *sync.WaitGroup = &sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(1)
	go add(100, 200, ch)
	result := <-ch //read the data from the channel
	wg.Wait()
	fmt.Println(result)
}

func add(x, y int, ch chan int) {

	result := x + y
	ch <- result //write data into the channel
	wg.Done()
}
