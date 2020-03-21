package main

import "fmt"

func main() {
	// student1 := [3]float64{5, 6, 1}
	// student2 := [3]float64{9, 8, 4}

	// var sum float64
	// sum += student1[0] + student1[1] + student1[2]
	// sum += student2[0] + student2[1] + student2[2]

	// const N = float64(len(student1) + len(student2))
	// fmt.Println(sum / N)

	student := [2][3]float64{
		{5, 6, 1},
		{9, 8, 4},
	}

	var sum float64
	sum += student[0][0] + student[0][1] + student[0][2] + student[1][0] + student[1][1] + student[1][2]
	const N = float64(len(student) * len(student[0]))
	fmt.Println(sum / N)
}
