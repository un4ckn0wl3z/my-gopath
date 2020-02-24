package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var (
		n   int
		err error
	)

	if a := os.Args; len(a) != 2 {
		// only: a var
		fmt.Println("Give me a number.")
	} else if n, err = strconv.Atoi(a[1]); err != nil {
		// only: a, n, err var
		fmt.Printf("Cannot convert %q", a[1])
	} else {
		// all
		n *= 2
		fmt.Printf("%s * 2 = %d\n", a[1], n)
	}

	fmt.Printf("n is %d. <> - you've been shadowed :]\n", n)

}
