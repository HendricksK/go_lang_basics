package main

import "fmt"

func main() {
	fmt.Println("hello its me")

	foo()

	fmt.Println("I love listening to rain ASMR")

	for i := 0; i <= 100; i++ {
		if i%2 == 0 { //Only print even numbers
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I am foo")
}

func bar() {
	fmt.Println("and then we exited")
}
