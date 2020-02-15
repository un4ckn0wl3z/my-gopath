package main

import "fmt"

func main() {
	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)

	fmt.Printf("Planet: %s\n", planet)
	fmt.Printf("Distance: %d\n", distance)
	fmt.Printf("Orbital: %f\n", orbital)
	fmt.Printf("Does %s Has Life: %t\n", planet, hasLife)
	fmt.Println()
	fmt.Println()
	fmt.Printf("Orbital: %f\n", orbital)
	fmt.Printf("Orbital: %.0f\n", orbital)
	fmt.Printf("Orbital: %.1f\n", orbital)
	fmt.Printf("Orbital: %.2f\n", orbital)
	fmt.Printf("Orbital: %.3f\n", orbital)

	/*
		%s = string
		%d = decimal
		%f = float
		%t = true or false
	*/

}
