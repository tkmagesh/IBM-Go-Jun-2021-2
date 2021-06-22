package main

import "fmt"

var operations = map[int]func(int, int) int{
	1: func(x, y int) int {
		return x + y
	},

	2: func(x, y int) int {
		return x - y
	},

	3: func(x, y int) int {
		return x * y
	},

	4: func(x, y int) int {
		return x / y
	},
}

func main() {
	for {
		choice := getUserChoice()
		if oper, exists := operations[choice]; exists {
			n1, n2 := getOperands()
			result := oper(n1, n2)
			fmt.Println(result)
			continue
		}
		break
	}
}

func getUserChoice() int {
	choice := 0
	fmt.Println("Enter the choice :")
	fmt.Println("1 : Add")
	fmt.Println("2 : Subtract")
	fmt.Println("3 : Multiply")
	fmt.Println("4 : Divide")
	fmt.Println("5 : Exit")
	fmt.Scanln(&choice)
	return choice
}

func getOperands() (int, int) {
	var n1, n2 int
	fmt.Println("Enter the operands")
	fmt.Scanf("%d %d\n", &n1, &n2)
	return n1, n2
}
