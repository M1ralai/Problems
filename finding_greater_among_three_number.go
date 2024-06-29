package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Println("Please give me 3 number, i will show you the greater one")
	fmt.Scanln(&a, &b, &c)
	if a > b {
		if a > c {
			d = a
		} else {
			d = c
		}
	} else {
		if b > c {
			d = b
		} else {
			d = c
		}
	}
	fmt.Printf("Greater number is: %v", d)
}
