package main

import (
	"fmt"
	"syscall"
)

func main() {

	str := "hello world"
	slice, _ := syscall.ByteSliceFromString(str)
	fmt.Printf("Type: %T \n", slice)
	fmt.Println("Bytes: ", slice)
	fmt.Println("String: ", string(slice))
}
