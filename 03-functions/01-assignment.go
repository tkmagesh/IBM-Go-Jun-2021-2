package main

import "fmt"

func main() {
	for {
		choice := getUserChoice()
		if choice < 1 || choice > 4 {
			break
		}
		n1, n2 := getOperands()
		switch choice {
		case 1:
			fmt.Println(add(n1, n2))
		case 2:
			fmt.Println(subtract(n1, n2))
		case 3:
			fmt.Println(multiply(n1, n2))
		case 4:
			fmt.Println(divide(n1, n2))
		}
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

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}
