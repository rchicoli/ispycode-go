package main

import (
	"fmt"
	"syscall"
)

func main() {

	var buffer = make([]byte, 128)
	fd, err := syscall.Open("/tmp/myfile.txt", syscall.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("Err:", err)
	}
	for {
		i, _ := syscall.Read(fd, buffer)
		if i <= 0 {
			break
		}
		fmt.Print(string(buffer[:i]))
	}
}
