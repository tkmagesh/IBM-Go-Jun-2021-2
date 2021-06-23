package main

import (
	"fmt"
	"sync"
	"time"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}

func main() {

	wg.Add(1)
	go print("Hello", wg, 1)

	wg.Add(1)
	go print("World", wg, 2)

	wg.Wait()
	fmt.Println("exiting main!")
}

func print(msg string, wg *sync.WaitGroup, goroutineId int) {
	fmt.Println("goroutine started => ", goroutineId)
	time.Sleep(3 * time.Second)
	fmt.Println(msg)
	wg.Done()
	fmt.Println("goroutine completed => ", goroutineId)
}
