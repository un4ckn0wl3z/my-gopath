package main

import "fmt"

func main() {
	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)

	fmt.Printf("Planet: %v\n", planet)
	fmt.Printf("Distance: %v\n", distance)
	fmt.Printf("Orbital: %v\n", orbital)
	fmt.Printf("Does %v Has Life: %v\n", planet, hasLife)

}
