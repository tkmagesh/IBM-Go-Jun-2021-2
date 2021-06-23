/* divide & conquer */
package main

import "fmt"

func sum(resultCh chan int, nos ...int) {
	result := 0
	for _, no := range nos {
		result += no
	}
	resultCh <- result
}

func main() {
	data := []int{4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55}
	//split the data into two sets
	firstSet := data[:len(data)/2]
	secondSet := data[len(data)/2:]

	//add those two sets of data concurrently
	ch := make(chan int)
	go sum(ch, firstSet...)
	go sum(ch, secondSet...)

	//get the results and add them to get the final total
	result := <-ch + <-ch

	//display the final total
	fmt.Println(result)
}
