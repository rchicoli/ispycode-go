package main

import (
	"fmt"
	"syscall"
)

func main() {

	id := syscall.Getgid()
	fmt.Println("Group ID:", id)
}
