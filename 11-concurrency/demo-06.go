package main

import (
	"fmt"
	"sync"
)

//"share memory by communicating (by using channels)"

var wg *sync.WaitGroup = &sync.WaitGroup{}

func main() {

	ch := make(chan int)
	//wg.Add(2)
	go add(100, 200, ch)
	go add(10, 20, ch)
	result := <-ch + <-ch
	//wg.Wait()
	//result := <-ch + <-ch //read the data from the channel - This line work fine becoz the chan is a buffered channel
	fmt.Println(result)
}

func add(x, y int, ch chan int) {

	result := x + y
	ch <- result //write data into the channel
	//wg.Done()
}
