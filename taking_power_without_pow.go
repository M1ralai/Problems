package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Give me two number")
	fmt.Scanln(&a, &b)
	constant := a
	fmt.Printf("I will take %v numbers %v degree power.", a, b)
	for i := 1; i < b; i++ {
		a *= constant
	}
	fmt.Printf("\nResult is: %v", a)
}
