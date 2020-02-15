package main

import "fmt"

func main() {
	celsius := 35.
	fahrenheit := (9*celsius + 160) / 5
	fmt.Printf("%g C is %g F", celsius, fahrenheit)
}
