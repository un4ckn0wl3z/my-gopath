package main

import "fmt"

/*
Duration type
custom type from Anuwat
*/
type Duration int64

func main() {

	var ns Duration
	var ms int64 = 1000

	ns = Duration(ms)
	ms = int64(ns)

	fmt.Println(ns)
	fmt.Println(ms)

}
