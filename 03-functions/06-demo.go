//deferred functions
package main

import "fmt"

func main() {
	defer fmt.Println("Exiting from main!")
	f1()
}

func f1() {
	var msg string
	defer func() {
		fmt.Println("[@defer-11] fn executed")
		fmt.Println(msg)
	}()
	defer fmt.Println("[@defer-12] fn executed")
	defer fmt.Println("[@defer-13] fn executed")
	fmt.Println("f1 invoked")
	msg = f2()

}

func f2() string {
	defer fmt.Println("[@defer-21] fn executed")
	defer fmt.Println("[@defer-22] fn executed")
	defer fmt.Println("[@defer-23] fn executed")
	fmt.Println("f2 invoked")
	return "Message from f2"
}
