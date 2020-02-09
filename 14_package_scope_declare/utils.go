package main

import "runtime"

func getGoVersion() string {
	return runtime.Version()
}
