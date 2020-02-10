package main

import "fmt"

func main() {
	speed := 100
	force := 2.5
	color := 65

	speed = int(float64(speed) * force)
	fmt.Println(speed)

	fmt.Println(string(color))

	fmt.Println(string([]byte{104, 105}))

	area := 10.5
	fmt.Println(area / 2)

}
