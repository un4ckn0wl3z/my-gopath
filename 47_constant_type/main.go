package main

import "fmt"

func main() {
	// const min int = 1
	// const pi float64 = 3.14
	// const version string = "2.0.1"
	// const debug bool = true

	const min = 1 + 2
	const pi = 3.14 * min
	const version = "2.0.1" + "-beta"
	const debug = !true

	fmt.Println(min, pi, version, debug)

}
