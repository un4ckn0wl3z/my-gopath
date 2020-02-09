package getcpunumber

import (
	"fmt"
	"runtime"
)

/*
PrintCPUNumber function
Get the CPU number of your system.
*/
func PrintCPUNumber() {
	fmt.Println(runtime.NumCPU())
}
