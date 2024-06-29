package main

import "fmt"

func main() {
	var n1, n2, ebob, big, small int
	fmt.Println("Bana iki tane sayı ver ver")
	fmt.Scanln(&n1, &n2)
	if n1 > n2 {
		big = n1
		small = n2
	} else {
		big = n2
		small = n1
	}
	for i := 2; i <= small; i++ {
		var a, b int
		a = small % i
		b = big % i
		if a == b {
			if a == 0 {
				ebob = i
				fmt.Println(i, "bu iki sayının ortak bölenidir")
			}
		} else {
			fmt.Println(i, "ortak bölen değil.")
		}
	}
	fmt.Println("EBOB: ", ebob)
}
