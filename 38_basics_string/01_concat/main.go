package main

import (
	"fmt"
	"strconv"
)

func main() {
	name, last := "Anuwat", "Khongchuai"
	name += " Palm"
	fmt.Println(name + " " + last)

	eq := "1 + 2 = "
	sum := 1 + 2
	fmt.Println(eq + strconv.Itoa(sum))

	eq = strconv.FormatBool(true) + " " + strconv.FormatBool(false)
	fmt.Println(eq)

}
