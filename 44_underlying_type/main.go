package main

import (
	"fmt"

	"github.com/un4ckn0wl3z/my-gopath/44_underlying_type/weights"
)

func main() {
	type (
		// Gram type is for me
		Gram int
		// Kilogram type is for me
		Kilogram Gram
		// Ton type is for me
		Ton Kilogram
	)

	var (
		salt   Gram     = 100
		apples Kilogram = 5
		truck  Ton      = 10
	)

	fmt.Printf("salt: %d, apples: %d, truck: %d\n", salt, apples, truck)

	salt = Gram(apples)
	apples = Kilogram(truck)
	truck = Ton(apples)

	salt = Gram(weights.Gram(100))
	fmt.Printf("Type of weights.Gram: %T\n", weights.Gram(100))
	fmt.Printf("Type of weights.Gram: %T\n", Gram(100))

	type myGram weights.Gram

	var paper myGram = 20
	fmt.Printf("Type of paper: %T\n", paper)

}
