package main

import "fmt"

func main() {
	var (
		planet   = "venus"
		distance = 261
	)

	fmt.Printf("%v is %v away. Think! %[2]v kms! %[1]v OMG!\n", planet, distance)

}
