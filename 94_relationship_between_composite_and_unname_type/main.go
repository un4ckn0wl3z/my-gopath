package main

import "fmt"

func main() {

	type bookcase [5]int
	type cabinet [5]int

	blue := bookcase{6, 9, 3, 2, 1}
	red := cabinet{6, 9, 3, 2, 1}

	fmt.Println("Are there equal? ")
	if blue == bookcase(red) {
		fmt.Println("v/")
	} else {
		fmt.Println("x")
	}

	fmt.Printf("blue: %#v\n", blue)
	fmt.Printf("red: %#v\n", red)

}
