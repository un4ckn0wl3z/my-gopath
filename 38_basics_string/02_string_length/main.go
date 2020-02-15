package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "Anuwat"
	fmt.Println(len(name))

	name = "อนุวัฒน์"
	fmt.Println(len(name))

	fmt.Println(utf8.RuneCountInString(name))

}
