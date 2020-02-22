package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const usage = `
Feet to Meters
-------------------------
This program converts feet to meters.

Usage:
feet [feetsToConvert]
`

func main() {

	if len(os.Args) < 2 {
		fmt.Printf(strings.TrimSpace(usage))
		return
	}

	arg := os.Args[1]
	feet, _ := strconv.ParseFloat(arg, 64)
	meters := feet * 0.3048

	fmt.Printf(" %g feet is %g meters.\n", feet, meters)

}
