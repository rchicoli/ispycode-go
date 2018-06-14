package main

import (
	"fmt"
	"syscall"
)

func main() {

	// get current working directory
	cwd, _ := syscall.Getwd()
	fmt.Println("cwd:", cwd)
}
