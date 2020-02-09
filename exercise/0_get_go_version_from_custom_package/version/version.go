package version

import "runtime"

/*
Version function
Get curernt Go version in your system.
*/
func Version() string {
	return runtime.Version()
}
