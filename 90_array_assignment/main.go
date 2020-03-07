package main

import "fmt"

func main() {

	blue := [3]int{6, 9, 3}

	red := blue
	red[2] = 50

	fmt.Printf("blue bookcase: %v\n", blue)
	fmt.Printf("red bookcase: %v\n", red)

}
