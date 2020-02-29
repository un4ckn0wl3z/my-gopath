package main

import "fmt"

func main() {

	switch {
	case 5 > 2:
		fmt.Println("5 > 2")
	case 3 < 1:
		fmt.Println("3 < 1")
	case 7 < 10:
		fmt.Println("7 < 10")
	default:
		fmt.Println("Meh?")
	}
}
