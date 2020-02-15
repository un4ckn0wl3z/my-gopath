package main

import "fmt"

func main() {
	var (
		myAge   = 30
		yourAge = 35
		average float64
	)

	average = float64(myAge+yourAge) / 2
	fmt.Println(average)

	ratio := 1.0 / 10
	fmt.Println(ratio)

	fmt.Println("sum:", 3+2)    // sum int
	fmt.Println("sum:", 2+3.5)  // sum float64
	fmt.Println("dif:", 3-1)    // diffence int
	fmt.Println("dif:", 3-0.5)  // diffence float64
	fmt.Println("prod:", 4*5)   // product int
	fmt.Println("prod:", 5*2.5) // product float64
	fmt.Println("quot:", 8/2)   // qoutient int
	fmt.Println("quot:", 8/1.5) // qoutient float64
	fmt.Println("rem:", 8%3)    // remainder only for int
	// fmt.Println("rem:", 8%0.3) -> error

	var degree float64 = 3.0 / 2
	fmt.Println(degree)

}
