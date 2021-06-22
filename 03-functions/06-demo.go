package main

import "fmt"

func main() {
	defer fmt.Println("Exiting from main!")
	f1()
}

func f1() {
	defer fmt.Println("[@defer-11] fn executed")
	defer fmt.Println("[@defer-12] fn executed")
	defer fmt.Println("[@defer-13] fn executed")
	fmt.Println("f1 invoked")
	f2()
}

func f2() {
	defer fmt.Println("[@defer-21] fn executed")
	defer fmt.Println("[@defer-22] fn executed")
	defer fmt.Println("[@defer-23] fn executed")
	fmt.Println("f2 invoked")
}
