package main

import "fmt"

func main() {

	const meters int = 100

	cm := 100
	m := cm / meters
	fmt.Printf("%dcm is %dm\n", cm, m)

	cm = 200
	m = cm / meters
	fmt.Printf("%dcm is %dm\n", cm, m)

	// total, numberOf := 5, 0
	// fmt.Println(total / numberOf) // runtime error

	// const total int = 5
	// const numberOf int = 0
	// fmt.Println(total / numberOf) // compile-time error

}
