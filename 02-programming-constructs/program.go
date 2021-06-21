package main

import "fmt"

func main() {
	//if else construct
	/*
		no := 21
		if no%2 == 0 {
			fmt.Printf("%d is an even number \n", no)
		} else {
			fmt.Printf("%d is an odd number \n", no)
		}
	*/

	if no := 21; no%2 == 0 {
		fmt.Printf("%d is an even number \n", no)
	} else {
		fmt.Printf("%d is an odd number \n", no)
	}

	//loop construct
	// for construct
	sum := 0

	for no := 1; no <= 10; no++ {
		sum += no
	}
	fmt.Printf("Sum of 1 - 10 integers = %d\n", sum)

	// 'for' used like a 'while' construct
	numSum := 1
	for numSum < 100 {
		numSum += numSum
	}
	fmt.Printf("numSum = %d\n", numSum)

	// 'for' for infinite loops
	x := 1
	for {
		x += x
		if x > 100 {
			break
		}
	}
	fmt.Printf("x = %d\n", x)

	// switch case
	newNo := 22
	/*
		switch newNo % 2 {
		case 0:
			fmt.Printf("%d is an even number\n", newNo)
		case 1:
			fmt.Printf("%d is an odd number\n", newNo)
		}
	*/

	switch {
	case newNo%2 == 0:
		fmt.Printf("%d is an even number\n", newNo)
	case newNo%2 == 1:
		fmt.Printf("%d is an odd number\n", newNo)
	}

	/*
		Rank by score
		score 0 - 3 => "Terrible"
		score 4 - 7 => "Not Bad"
		score 8 - 9 => "Good"
		score 10 => "Excellent"
		otherwise => "Unknown score"
	*/
	score := 6
	/*
		switch score {
		case 0:
		case 1:
		case 2:
		case 3:
			fmt.Println("Terrible")
		case 4:
		case 5:
		case 6:
		case 7:
			fmt.Println("Not bad!")
		case 8:
		case 9:
			fmt.Println("Good")
		case 10:
			fmt.Println("Excellent")
		default:
			fmt.Println("Unknown score")
		}
	*/

	/*
		switch score {
		case 0, 1, 2, 3:
			fmt.Println("Terrible")
		case 4, 5, 6, 7:
			fmt.Println("Not bad!")
		case 8, 9:
			fmt.Println("Good")
		case 10:
			fmt.Println("Excellent")
		default:
			fmt.Println("Unknown score")
		}
	*/
	switch {
	case score >= 0 && score <= 3:
		fmt.Println("Terrible")
	case score >= 4 && score <= 7:
		fmt.Println("Not bad!")
	case score == 8 || score == 9:
		fmt.Println("Good")
	case score == 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Unknown score")
	}

	//using fallthrough
	n := 4
	switch n {
	case 0:
		fmt.Println("is zero")
	case 1:
		fmt.Println("n <= 1")
		fallthrough
	case 2:
		fmt.Println("n <= 2")
		fallthrough
	case 3:
		fmt.Println("n <= 3")
		fallthrough
	case 4:
		fmt.Println("n <= 4")
		fallthrough
	case 5:
		fmt.Println("n <= 5")
		fallthrough
	case 6:
		fmt.Println("n <= 6")
		fallthrough
	case 7:
		fmt.Println("n <= 7")
	case 8:
		fmt.Println("n <= 8")
		fallthrough
	case 9:
		fmt.Println("n <= 9")
		fallthrough
	case 10:
		fmt.Println("n <= 10")
	}
}
