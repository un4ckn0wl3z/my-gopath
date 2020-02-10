package main

import (
	"fmt"
	"path"
)

func main() {
	var file string
	_, file = path.Split("secret/key.txt")

	fmt.Println("file: ", file)

}
