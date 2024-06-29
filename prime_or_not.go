package main

import "fmt"

func main() {
	var number, a int
	fmt.Println("Write me a number then i will say it is prime or not")
	fmt.Scanln(&number)
	for i := 2; i < number/2; i++ {
		a = number % i
		if a == 0 {
			fmt.Println("This number is not prime")
			return
		}
	}
	if a != 0 {
		fmt.Println("This number is prime")
	}
}
