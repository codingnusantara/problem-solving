package main

import (
	"fmt"
)

// Problem :

// Make a function with the following output :
// Output code :

// xoxox
// oxoxo
// xoxox

func main() {
	var x, y int

	fmt.Print("Enter x: ")
	fmt.Scan(&x)

	fmt.Print("Enter y: ")
	fmt.Scan(&y)

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if (i+j)%2 == 0 {
				fmt.Print("x")
			} else {
				fmt.Print("o")
			}
		}
		fmt.Println()
	}
}
