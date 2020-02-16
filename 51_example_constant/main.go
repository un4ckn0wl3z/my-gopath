package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Second * 10
	fmt.Println(t)

	const i = 10
	t = time.Second * i
	fmt.Println(t)

}
