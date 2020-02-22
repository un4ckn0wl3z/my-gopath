package main

import "fmt"

func main() {
	score, valid := 2, true
	if score > 3 && valid {
		fmt.Println("Good!")
	} else if score == 3 {
		fmt.Println("on the edge!")
	} else if score == 2 {
		fmt.Println("meh..!")
	} else {
		fmt.Println("low!")

	}
}
