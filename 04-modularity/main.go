package main

import (
	"fmt"
	calc "myApp/calculator"
	"myApp/calculator/utils"
	"strconv"

	"github.com/fatih/color"
)

func main() {
	//fmt.Println(calc.AddClient(10, 20))
	add := calc.GETADDER()
	color.Green(strconv.Itoa(add(100, 200)))
	fmt.Println(calc.Subtract(100, 200))
	fmt.Println(utils.Multiply(10, 20))
	fmt.Println(utils.Divide(100, 7))
}
