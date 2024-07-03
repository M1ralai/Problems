package main

import (
	"fmt"
)

func main() {
	var imput string
	var counter [26]int
	fmt.Println("Give me a word")
	fmt.Scanln(&imput)
	for i := 0; i < len(imput); i++ {
		fmt.Println(string(imput[i]))
		for b := 0; b <= 26; b++ {
			for c := 97; c+b == int(imput[i]); c++ {
				counter[b] += 1
			}
		}
	}
	for b := 0; b <= 25; b++ {
		if counter[b] != 0 {
			fmt.Println("This word have ", counter[b], "times", string(b+97))
		}
	}
}
