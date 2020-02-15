package main

import "fmt"

func main() {
	fmt.Printf("%b\n", 0)
	fmt.Printf("%b\n", 1)

	fmt.Printf("%02b\n", 0)
	fmt.Printf("%02b\n", 1)
	fmt.Printf("%02b\n", 2)
	fmt.Printf("%02b\n", 3)

	fmt.Printf("%08b\n", 255)

}
