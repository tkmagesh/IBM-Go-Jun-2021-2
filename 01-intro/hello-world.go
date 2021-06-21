/* package declaration */
package main

import "fmt"

/* import packages */

/* package level variables */
var name, msg = "Magesh", "Have a nice day!"

/* package init function */

/* main function */
func main() {
	/*
		var msg string
		msg = "Hello World"
	*/

	/*
		var msg string = "Hello World"
	*/

	/*
		var msg = "Hello World"
	*/
	//the following syntax is applicable only in a function (not in a package level variable)
	/*
		msg := "Hello World"
	*/
	//fmt.Println(msg)

	//declaraing multiple variables
	/*
		var name string
		var msg string
	*/
	/*
		var (
			name string
			msg  string
		)
	*/
	/*
		var (
			name, msg string
		)
	*/
	/*
		var name, msg string
		name = "Magesh"
		msg = "Have a nice day!"
	*/
	/*
		var name, msg string = "Magesh", "Have a nice day!"
	*/
	/*
		var name, msg = "Magesh", "Have a nice day!"
	*/
	/*
		name, msg := "Magesh", "Have a nice day!"
	*/
	//fmt.Println(name, msg)
	fmt.Printf("Hi %s, %s\n", name, msg)
}

/* other functions */
