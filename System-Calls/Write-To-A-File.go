package main

import (
	"fmt"
	"syscall"
)

func main() {

	// string to write to file
	bytes := []byte("The cat in the hat\n")

	// open file for writing
	fd1, err := syscall.Open("/tmp/myfile.txt", syscall.O_CREAT|syscall.O_RDWR|syscall.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Err:", err)
	}

	// write to file
	fd2, err := syscall.Write(fd1, bytes)

	// check results
	if err == nil {
		fmt.Println("Successfully wrote to fd:", fd2)
	} else {
		fmt.Println("Err:", err)

	}
}
