package main

import (
	"fmt"
	"syscall"
)

func main() {
	err := syscall.Mkdir("/tmp/mydir", 0754)

	if err == nil {
		fmt.Println("success")
	}
}
