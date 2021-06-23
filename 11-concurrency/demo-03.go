package main

import (
	"fmt"
	"sync"
)

//communicate by sharing memory - DONT DO THIS (instead try "share memory by communicating (by using channels)")
var result int

var wg *sync.WaitGroup = &sync.WaitGroup{}
var mutex = &sync.Mutex{}

func main() {
	wg.Add(1)
	go add(100, 200)
	wg.Wait()
	fmt.Println(result)
}

func add(x, y int) {
	mutex.Lock()
	result = x + y
	mutex.Unlock()
	wg.Done()
}
