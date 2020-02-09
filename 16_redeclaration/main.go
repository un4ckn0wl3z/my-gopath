package main

import "fmt"

func main() {
	var safe bool
	safe, speed := true, 50

	var name string = "Commando"
	name = "Polisize"

	fmt.Println(safe, speed)
	fmt.Println(name)

	personName := "Nikola"
	personName, age := "Robert", 46
	fmt.Println(personName, age)

}
