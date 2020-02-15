package main

import "fmt"

func main() {
	width, height := 5., 12.
	area := width * height

	fmt.Printf("aread=%g\n", area)

	area = area - 10
	area = area + 10
	area = area * 2
	area = area / 2
	fmt.Printf("aread=%g\n", area)

	area -= 10
	area += 10
	area *= 10
	area /= 10
	fmt.Printf("aread=%g\n", area)
}
