package main

import "fmt"

func main() {
	var number int
	var prime_indicator, loop bool
	prime_indicator = false
	for !loop {
		fmt.Println("Hello give me a number then  will write prime number until this number, 0 is quit")
		fmt.Scanln(&number)
		if number == 0 {
			loop = true
		} else if number <= 1 {
			fmt.Println("Please write number greater than 2")
		} else {
			for i := 2; i <= number; i++ {
				prime_indicator = true
				for a := 2; a <= i/2; a++ {
					mod := i % a
					if mod == 0 {
						prime_indicator = false
					}
				}
				if prime_indicator == true {
					fmt.Println(i)
				}
			}
		}
	}
}
