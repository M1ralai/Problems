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
		for b := 0; b <= 26; b++ {
			for c := 97; c+b == int(imput[i]); c++ {
				counter[b] += 1
			}
		}
	}
	fmt.Println("This Word Have: ", counter[0], " a letter")
	fmt.Println("This Word Have: ", counter[1], " b letter")
	fmt.Println("This Word Have: ", counter[2], " c letter")
	fmt.Println("This Word Have: ", counter[3], " d letter")
	fmt.Println("This Word Have: ", counter[4], " e letter")
	fmt.Println("This Word Have: ", counter[5], " f letter")
	fmt.Println("This Word Have: ", counter[6], " g letter")
	fmt.Println("This Word Have: ", counter[7], " h letter")
	fmt.Println("This Word Have: ", counter[8], " ı letter")
	fmt.Println("This Word Have: ", counter[9], " j letter")
	fmt.Println("This Word Have: ", counter[10], " k letter")
	fmt.Println("This Word Have: ", counter[11], " l letter")
	fmt.Println("This Word Have: ", counter[12], " m letter")
	fmt.Println("This Word Have: ", counter[13], " n letter")
	fmt.Println("This Word Have: ", counter[14], " o letter")
	fmt.Println("This Word Have: ", counter[15], " p letter")
	fmt.Println("This Word Have: ", counter[16], " q letter")
	fmt.Println("This Word Have: ", counter[17], " r letter")
	fmt.Println("This Word Have: ", counter[18], " s letter")
	fmt.Println("This Word Have: ", counter[19], " t letter")
	fmt.Println("This Word Have: ", counter[20], " u letter")
	fmt.Println("This Word Have: ", counter[21], " v letter")
	fmt.Println("This Word Have: ", counter[22], " w letter")
	fmt.Println("This Word Have: ", counter[23], " x letter")
	fmt.Println("This Word Have: ", counter[24], " y letter")
	fmt.Println("This Word Have: ", counter[25], " z letter")
}
