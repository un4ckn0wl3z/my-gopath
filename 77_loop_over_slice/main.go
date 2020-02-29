package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%q\n", os.Args[i])
	}

	words := strings.Fields(
		"I cut my hair becuz you dont care my heart",
	)

	for i := 0; i < len(words); i++ {
		fmt.Printf("%q\n", words[i])
	}
}
