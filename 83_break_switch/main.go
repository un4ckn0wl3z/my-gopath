package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "I cut my hair becuz you dont care my heart hair"

func main() {
	words := strings.Fields(corpus)
	query := os.Args[1:]
queries:
	for _, q := range query {
	search:
		for i, w := range words {

			switch q {
			case "cut", "heart":
				break search
			}

			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				continue queries
			}
		}
	}
}
