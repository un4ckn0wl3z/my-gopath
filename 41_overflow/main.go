package main

import "fmt"

func main() {
	var (
		width  uint8 = 255
		height       = 255
	)

	width++

	if int(width) < height {
		fmt.Println("height is greater width")
	}
}
