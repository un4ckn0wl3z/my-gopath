package main

import (
	"fmt"
	"path"
)

func main() {
	var dir, file string

	dir, file = path.Split("secret/key.txt")

	fmt.Println("dir: ", dir)
	fmt.Println("file: ", file)

}
