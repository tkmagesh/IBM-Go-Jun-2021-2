package main

import "fmt"

func main() {
	var no int = 10
	var noPtr *int = &no
	fmt.Println(no, noPtr)

	//deferencing a pointer
	value := *noPtr
	fmt.Println(value)

	/* value = *(&value) */
	fmt.Println("Before incrementing, no = ", no)
	increment(&no)
	fmt.Println("after incrementing, no = ", no)

	x, y := 10, 20
	fmt.Println("Before swaping", x, y)
	swap(&x, &y)
	fmt.Println("After swaping", x, y)

	nos := []int{10, 20, 30}
	addValue(&nos, 40)
	fmt.Println(nos) //=> should print 10, 20, 30, 40
}

func addValue(list *[]int, value int) {
	*list = append(*list, value)
}

func increment(x *int) {
	*x += 1
}

func swap(x, y *int) {
	*x, *y = *y, *x
}
