package main

import "fmt"

func main() {
	var speed int
	fmt.Println(speed)

	speed = 100
	fmt.Println(speed)

	speed = speed - 25
	fmt.Println(speed)

	var (
		name   string
		age    int
		famous bool
	)

	name = "Newton"
	age = 84
	famous = true
	fmt.Println(name, age, famous)

	name = "Somebody"
	age = 20
	famous = false
	fmt.Println(name, age, famous)

	prevName := name

	name = "Robert"

	fmt.Println("Previous name:", prevName)
	fmt.Println("Current name:", name)

}
