package main

import "fmt"

var version string = "1.0.0"

func main() {
	// normal
	var (
		firstname string = "Anuwat"
		lastname  string = "Kh."
		age       int    = 24
	)

	// score := 0 // dont

	var score int // already 0
	fmt.Println(score)
	fmt.Println(firstname, lastname, age)

	// short

	height, width := 150, 240

	width, color := 300, "red"

	fmt.Println(height, width, color)

}

/*
Use Declaration
1. When you dont know initial value.
2. When you need a package scope variable.
3. When you wanna group variables together for greater readability
*/

/*
Use Short Declaration
1. When you know the initial value
2. To keep code concise
3. For redeclaration
*/
