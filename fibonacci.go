package main

import "fmt"

func main() {
	var c int
	a := 1
	b := 1
	d := 1
	fmt.Println("Please give me a number so i could write fibonacci array until this step.")
	fmt.Scanln(&c)
	for i := 1; i <= c; i++ {
		fmt.Printf("%v. %v\n", i, d)
		d = a + b
		a = b
		b = d
	}
}
