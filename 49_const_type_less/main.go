package main

import "fmt"

func main() {
	const min = 42
	var f float64
	f = min

	fmt.Println(f)

	const max = 42

	var i int = max
	var j float64 = max
	var b byte = max

	fmt.Println(i, j, b)

}
