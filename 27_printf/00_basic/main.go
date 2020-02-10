package main

import "fmt"

func main() {
	var brand = "google"
	fmt.Printf("%q\n", brand)
	fmt.Printf("%s\n", brand)

	ops := 2350
	ok := 543
	fail := 443

	fmt.Println(
		"total:", ops, "success:", ok, "/", fail,
	)

	fmt.Printf(
		"total: %d success: %d / %d  \n", ops, ok, fail,
	)

}
