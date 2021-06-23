package main

import (
	"fmt"
	"sync"
	"time"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}

func main() {

	wg.Add(1)
	go print("Hello", wg)

	wg.Add(1)
	go print("World", wg)

	wg.Wait()
	fmt.Println("exiting main!")
}

func print(msg string, wg *sync.WaitGroup) {
	time.Sleep(3 * time.Second)
	fmt.Println(msg)
	wg.Done()
}
