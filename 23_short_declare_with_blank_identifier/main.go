package main

import (
	"fmt"
	"path"
)

func main() {
	_, file := path.Split("secret/key.txt")
	fmt.Println("file: ", file)
}
