package main

import "fmt"

func main() {
	color := "red"
	fmt.Println("reddish color?", color == "red" || color == "dark red")

	color = "dark red"
	fmt.Println("reddish color?", color == "red" || color == "dark red")

	color = "green"
	fmt.Println("reddish color?", color == "red" || color == "dark red")
}
