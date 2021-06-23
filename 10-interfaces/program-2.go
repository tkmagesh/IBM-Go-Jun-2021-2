package main

import "fmt"

type Employee struct {
	Name string
}

func main() {
	var x interface{}
	x = 10
	x = "Magesh"
	x = true

	x = 100
	if no, ok := x.(int); ok {
		fmt.Println("Double of x is ", no*2)
	} else {
		fmt.Println("x cannot be treated as an int")
	}
	fmt.Println(x)

	//x = 200
	//x = "Magesh"
	//x = true
	x = Employee{Name: "Suresh"}

	switch val := x.(type) {
	case int:
		fmt.Println("Double of x is ", val*2)
	case string:
		fmt.Println("Len of x is ", len(val))
	case bool:
		fmt.Println("x is boolean value = ", val)
	case Employee:
		fmt.Println("x is an employee with name = ", val.Name)
	default:
		fmt.Println("Unknown type")
	}
}
