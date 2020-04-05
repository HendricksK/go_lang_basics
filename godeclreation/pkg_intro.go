package main

import (
	"fmt"
)

var y = 43 // Declaring with var allowed outside of body. Scope
var z int  // Declreation of var z of type int, but no value assigned
// Each element of such a variable or value is set to the zero value for its type: false for booleans, 0 for numeric types, "" for strings, and nil for pointers, functions, interfaces, slices, channels, and maps https://golang.org/ref/spec#The_zero_value

func main() {
	x := 42 // Declares variable x and assigns a value
	n, err := fmt.Println("Hello, playground", 239489234, true)
	if err != nil { // Checking for nill
		fmt.Println(err)
	}
	fmt.Println("value of n", n)
	fmt.Println("value of x", x)
	x = 99
	fmt.Println("value of x", x)

	foo()
}

func foo() {
	// Declaring a varibale and assign the value using the short declreation operator, operating on opperands
	// Declaring of a certain type
	// Declaring of var outside of function body allows us to use variables that have been declared outside of function body
	fmt.Println("Variable y decalred as var", y)
	fmt.Println("Declreation of z with no value but type int = ", z)
}
