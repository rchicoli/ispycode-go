package main

import (
	"fmt"
	"syscall"
)

func main() {

	// change working directory
	syscall.Chdir("/home/dennis/ISPY")

	// get current working directory
	cwd, _ := syscall.Getwd()
	fmt.Println("cwd:", cwd)
}
