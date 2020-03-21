package main

import "fmt"

func main() {

	student := [2][3]float64{
		{5, 6, 1},
		{9, 8, 4},
	}

	var sum float64

	for _, grades := range student {
		for _, grade := range grades {
			sum += grade
		}
	}

	const N = float64(len(student) * len(student[0]))
	fmt.Println(sum / N)
}
