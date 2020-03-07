package main

import "fmt"

func main() {
	books := [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Editon",
	}

	for _, v := range books {
		fmt.Println(v)
	}

}
