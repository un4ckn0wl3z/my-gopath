package main

import (
	"fmt"
	"os"
)

func main() {
	// for i := 1; i < len(os.Args); i++ {
	// 	fmt.Printf("%q\n", os.Args[i])
	// }

	for _, v := range os.Args[1:] {
		fmt.Printf("%q\n", v)
	}
}
