package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
)

func main() {

	// first arg is the process id
	arg0 := os.Args[0]
	val0, _ := strconv.ParseInt(arg0, 10, 32)
	pid := int(val0)

	// second arg is the signal number
	arg1 := os.Args[1]
	val1, _ := strconv.ParseInt(arg1, 10, 32)
	signal := int(val1)

	err := syscall.Kill(pid, signal)

	if err == nil {
		fmt.Println("success")
	}
}
