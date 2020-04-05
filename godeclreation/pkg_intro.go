package main

import (
	"fmt"
)

func main() {
	n, err := fmt.Println("Hello, playground", 239489234, true)
	if err != nil { // Checking for nill
		fmt.Println(err)
	}
	fmt.Println(n)
}
