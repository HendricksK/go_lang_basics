package main

import (
	"fmt"
)

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
	y := 100 + 34234 // This is a statement
	fmt.Println("value of y y := 100 + 34234", y)
}
